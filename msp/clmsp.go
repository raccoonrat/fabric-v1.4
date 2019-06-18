/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/bccsp"
	clbccsp "github.com/hyperledger/fabric/bccsp/cl"
	"github.com/hyperledger/fabric/bccsp/cl/signer"
	"github.com/hyperledger/fabric/bccsp/sw"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/protos/msp"
	m "github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

var clmspLogger = flogging.MustGetLogger("clmsp")

// This is an instantiation of an MSP that
// uses BCCSP for its cryptographic primitives.
type clmsp struct {
	// version specifies the behaviour of this msp
	version MSPVersion

	// list of KGC pubs we trust
	rootPubs []bccsp.Key

	// list of CA TLS certs we trust
	tlsRootCerts [][]byte

	// list of intermediate TLS certs we trust
	tlsIntermediateCerts [][]byte

	// list of signing identities
	signer SigningIdentity

	// list of admin identities
	admins []clidentity

	// the crypto provider
	csp bccsp.BCCSP

	// the provider identifier for this MSP
	name string

	// verification options for MSP members TLS
	opts *x509.VerifyOptions

	// list of certificate revocation lists
	CRL []*pkix.CertificateList

	// list of OUs
	ouIdentifiers map[string][][]byte

	// cryptoConfig contains
	cryptoConfig *m.FabricCryptoConfig

	// NodeOUs configuration
	ouEnforcement bool
	// These are the OUIdentifiers of the clients, peers and orderers.
	// They are used to tell apart these entities
	clientOU, peerOU *OUIdentifier
}

// newBccspMsp returns an MSP instance backed up by a BCCSP
// crypto provider. It handles x.509 certificates and can
// generate identities and signing identities backed by
// certificates and keypairs
func newIBPCLAMsp(version MSPVersion) (MSP, error) {
	mspLogger.Debugf("Creating IBPCLA-based MSP instance")

	csp, err := clbccsp.New(sw.NewDummyKeyStore())
	if err != nil {
		panic(fmt.Sprintf("unexpected condition, error received [%s]", err))
	}

	msp := &clmsp{}
	msp.version = version
	msp.csp = csp

	return msp, nil
}

func (msp *clmsp) getCertFromPem(idBytes []byte) (*x509.Certificate, error) {
	if idBytes == nil {
		return nil, errors.New("getCertFromPem error: nil idBytes")
	}

	// Decode the pem bytes
	pemCert, _ := pem.Decode(idBytes)
	if pemCert == nil {
		return nil, errors.Errorf("getCertFromPem error: could not decode pem bytes [%v]", idBytes)
	}

	// get a cert
	var cert *x509.Certificate
	cert, err := x509.ParseCertificate(pemCert.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "getCertFromPem error: failed to parse x509 cert")
	}

	return cert, nil
}

func (msp *clmsp) getclAdminIdentityFromConf(adminconfig *msp.CLMSPAdminConfig) (*clidentity, error) {

	// get the PA in the right format
	/*
		block, _ := pem.Decode(PABytes)
		if block == nil {
			return nil, errors.New("invalid PA, failed decoding pem Bytes")

		}
	*/
	KGCID, err := msp.rootPubs[0].Bytes()
	if err != nil {
		return nil, errors.WithMessage(err, "getAdminIdentityFromConf error: Failed to load KGCID")
	}

	// Use the hash of the identity's certificate as id in the IdentityIdentifier
	hashOpt, err := bccsp.GetHashOpt(msp.cryptoConfig.IdentityIdentifierHashFunction)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function options")
	}

	digest, err := msp.csp.Hash(adminconfig.PA, hashOpt)
	if err != nil {
		return nil, errors.WithMessage(err, "failed hashing PA to compute the id of the IdentityIdentifier")
	}
	//set OU and Role
	ou := &m.OrganizationUnit{
		MspIdentifier:                msp.name,
		OrganizationalUnitIdentifier: adminconfig.OU,
		CertifiersIdentifier:         KGCID,
	}
	if strings.ToUpper(adminconfig.Role) != m.MSPRole_ADMIN.String() {
		return nil, errors.New("failed generate admin identity from config, Role is not ADMIN")
	}
	role := &m.MSPRole{
		MspIdentifier: msp.name,
		Role:          m.MSPRole_ADMIN,
	}

	id := &IdentityIdentifier{
		Mspid: msp.name,
		Id:    hex.EncodeToString(digest)}

	return &clidentity{
		PA:     adminconfig.PA,
		id:     id,
		msp:    msp,
		nameID: adminconfig.ID,
		OU:     ou,
		Role:   role,
	}, nil
}

/*
func (msp *clmsp) getclIdentityFromConfPA(idBytes []byte) (Identity, bccsp.Key, error) {

	//get the PA in the right format
	PA, err := msp.getCertFromPem(idBytes)
	if err != nil {
		return nil, nil, err
	}

	//to do recover Pubs from PA
	PubK, err := msp.csp.KeyImport(idBytes, &bccsp.ECDSAPKIXPublicKeyImportOpts{Temporary: True})

	mspId, err := newclIdentity(PubK, msp)
	if err != nil {
		return nil, nil, err
	}

	return mspId, PubK, nil
}

func (msp *clmsp) getIdentityFromConf(idBytes []byte) (Identity, bccsp.Key, error) {
	// get a ca cert
	cert, err := msp.getCertFromPem(idBytes)
	if err != nil {
		return nil, nil, err
	}

	// get the public key in the right format
	certPubK, err := msp.csp.KeyImport(cert, &bccsp.X509PublicKeyImportOpts{Temporary: true})

	mspId, err := newIdentity(cert, certPubK, msp)
	if err != nil {
		return nil, nil, err
	}

	return mspId, certPubK, nil
}
*/

func (msp *clmsp) getSigningIdentityFromConf(sidInfo *m.CLMSPSignerConfig) (SigningIdentity, error) {
	if sidInfo == nil {
		return nil, errors.New("getIdentityFromBytes error: nil sidInfo")
	}
	//msp.getKGCIdentifier
	KGCID, err := msp.rootPubs[0].Bytes()
	if err != nil {
		return nil, errors.WithMessage(err, "getIdentityFromBytes error: Failed to load KGCID")
	}
	/*
			pemKey, _ := pem.Decode(sidInfo.Sk)
			if pemKey == nil {
				return nil, errors.New("getIdentityFromBytes error: Failed to load Sk")
			}
		privKey, err := msp.csp.KeyImport(pemKey.Bytes, &bccsp.CLPrivateKeyImportOpts{Temporary: true})
	*/
	privKey, err := msp.csp.KeyImport(sidInfo.Sk, &bccsp.CLPrivateKeyImportOpts{Temporary: true})
	if err != nil {
		return nil, errors.WithMessage(err, "getIdentityFromBytes error: Failed to import EC private key")
	}

	// get the peer signer
	peerSigner, err := signer.New(msp.csp, privKey)
	if err != nil {
		return nil, errors.WithMessage(err, "getIdentityFromBytes error: Failed initializing bccspCryptoSigner")
	}
	//set OU and Role
	ou := &m.OrganizationUnit{
		MspIdentifier:                msp.name,
		OrganizationalUnitIdentifier: sidInfo.OU,
		CertifiersIdentifier:         KGCID,
	}
	role := &m.MSPRole{
		MspIdentifier: msp.name,
		Role:          m.MSPRole_MEMBER,
	}
	//m.MSPRole_MEMBER,
	if strings.ToUpper(sidInfo.Role) == m.MSPRole_ADMIN.String() {
		role.Role = m.MSPRole_ADMIN
	}
	return newCLSigningIdentity(sidInfo.PA, sidInfo.ID, ou, role, peerSigner, msp)
}

// Setup sets up the internal data structures
// for this MSP, given an MSPConfig ref; it
// returns nil in case of success or an error otherwise
func (msp *clmsp) Setup(conf1 *m.MSPConfig) error {
	if conf1 == nil {
		return errors.New("Setup error: nil conf reference")
	}

	if conf1.Type != int32(IBPCLA) {
		return errors.Errorf("setup error: config is not of type IBPCLA")
	}

	// given that it's an msp of type fabric, extract the MSPConfig instance
	conf := &m.CLMSPConfig{}
	err := proto.Unmarshal(conf1.Config, conf)
	if err != nil {
		return errors.Wrap(err, "failed unmarshalling ibpcla msp config")
	}

	// set the name for this msp
	if len(conf.Name) == 0 {
		return errors.New("need a valid Name")
	}
	msp.name = conf.Name
	mspLogger.Debugf("Setting up IBPCLA MSP instance %s", msp.name)

	// setup
	err = msp.setupV11(conf)
	if err != nil {
		return errors.Wrap(err, "failed setup IBPCLA MSP")
	}
	return nil
}

// GetVersion returns the version of this MSP
func (msp *clmsp) GetVersion() MSPVersion {
	return msp.version
}

// GetType returns the type for this MSP
func (msp *clmsp) GetType() ProviderType {
	return IBPCLA
}

// GetIdentifier returns the MSP identifier for this instance
func (msp *clmsp) GetIdentifier() (string, error) {
	return msp.name, nil
}

// GetTLSRootCerts returns the root certificates for this MSP
func (msp *clmsp) GetTLSRootCerts() [][]byte {
	return msp.tlsRootCerts
}

// GetTLSIntermediateCerts returns the intermediate root certificates for this MSP
func (msp *clmsp) GetTLSIntermediateCerts() [][]byte {
	return msp.tlsIntermediateCerts
}

// GetDefaultSigningIdentity returns the
// default signing identity for this MSP (if any)
func (msp *clmsp) GetDefaultSigningIdentity() (SigningIdentity, error) {
	mspLogger.Debugf("Obtaining default IBPCLA signing identity")

	if msp.signer == nil {
		return nil, errors.New("this MSP does not possess a valid default signing identity")
	}

	return msp.signer, nil
}

// GetSigningIdentity returns a specific signing
// identity identified by the supplied identifier
func (msp *clmsp) GetSigningIdentity(identifier *IdentityIdentifier) (SigningIdentity, error) {
	// TODO
	return nil, errors.Errorf("no signing identity for %#v", identifier)
}

// Validate attempts to determine whether
// the supplied identity is valid according
// to this MSP's roots of trust; it returns
// nil in case the identity is valid or an
// error otherwise
func (msp *clmsp) Validate(id Identity) error {
	var clid *clidentity
	switch t := id.(type) {
	case *clidentity:
		clid = id.(*clidentity)
	case *clsigningidentity:
		clid = id.(*clsigningidentity).clidentity
	default:
		return errors.Errorf("identity type %T is not recognized", t)
	}

	mspLogger.Debugf("CLMSP %s validating identity", msp.name)
	if clid.GetMSPIdentifier() != msp.name {
		return errors.Errorf("the supplied identity does not belong to this msp")
	}
	return clid.validateIdentity()
}

// hasOURole checks that the identity belongs to the organizational unit
// associated to the specified MSPRole.
// This function does not check the certifiers identifier.
// Appropriate validation needs to be enforced before.
func (msp *clmsp) hasOURole(id Identity, mspRole m.MSPRole_MSPRoleType) error {
	// Check NodeOUs
	if !msp.ouEnforcement {
		return errors.New("NodeOUs not activated. Cannot tell apart identities.")
	}

	mspLogger.Debugf("MSP %s checking if the identity is a client", msp.name)

	switch id := id.(type) {
	// If this identity is of this specific type,
	// this is how I can validate it given the
	// root of trust this MSP has
	case *clidentity:
		return msp.hasOURoleInternal(id, mspRole)
	default:
		return errors.New("Identity type not recognized")
	}
}

func (msp *clmsp) hasOURoleInternal(id *clidentity, mspRole m.MSPRole_MSPRoleType) error {
	var nodeOUValue string
	switch mspRole {
	case m.MSPRole_CLIENT:
		nodeOUValue = msp.clientOU.OrganizationalUnitIdentifier
	case m.MSPRole_PEER:
		nodeOUValue = msp.peerOU.OrganizationalUnitIdentifier
	default:
		return fmt.Errorf("Invalid MSPRoleType. It must be CLIENT, PEER or ORDERER")
	}

	for _, OU := range id.GetOrganizationalUnits() {
		if OU.OrganizationalUnitIdentifier == nodeOUValue {
			return nil
		}
	}

	return nil
	//return fmt.Errorf("The identity does not contain OU [%s], MSP: [%s]", mspRole, msp.name)
}

// DeserializeIdentity returns an Identity given the byte-level
// representation of a SerializedIdentity struct
func (msp *clmsp) DeserializeIdentity(serializedID []byte) (Identity, error) {
	mspLogger.Debug("Obtaining identity")

	// We first deserialize to a SerializedIdentity to get the MSP ID
	sId := &m.SerializedIdentity{}
	err := proto.Unmarshal(serializedID, sId)
	if err != nil {
		return nil, errors.Wrap(err, "could not deserialize a SerializedIdentity")
	}

	if sId.Mspid != msp.name {
		return nil, errors.Errorf("expected MSP ID %s, received %s", msp.name, sId.Mspid)
	}
	mspLogger.Debug("Mspid:", sId.Mspid)

	return msp.deserializeIdentityInternal(sId.IdBytes)
}

// deserializeIdentityInternal returns an identity given its byte-level representation
func (msp *clmsp) deserializeIdentityInternal(serializedIdentity []byte) (Identity, error) {
	mspLogger.Debug("clmsp: deserializing identity")
	serialized := new(m.SerializedIBPCLAIdentity)
	err := proto.Unmarshal(serializedIdentity, serialized)
	if err != nil {
		return nil, errors.Wrap(err, "could not deserialize a Serialized CLIdentity")
	}
	if serialized.PA == nil {
		return nil, errors.Errorf("unable to deserialize CLIdentity: PA is invalid")
	}

	mspLogger.Debug("clmsp: deserializing identity", serialized.ID)

	//OU
	ou := &m.OrganizationUnit{}
	err = proto.Unmarshal(serialized.Ou, ou)
	if err != nil {
		return nil, errors.Wrap(err, "cannot deserialize the OU of the identity")
	}
	//Role
	role := &m.MSPRole{}
	err = proto.Unmarshal(serialized.Role, role)
	if err != nil {
		return nil, errors.Wrap(err, "cannot deserialize the role of the identity")
	}
	return newclIdentity(serialized.PA, msp, serialized.ID, ou, role)
}

// SatisfiesPrincipal returns null if the identity matches the principal or an error otherwise
func (msp *clmsp) SatisfiesPrincipal(id Identity, principal *m.MSPPrincipal) error {
	err := msp.Validate(id)
	if err != nil {
		return errors.Wrap(err, "identity is not valid with respect to this MSP")
	}

	return msp.satisfiesPrincipalValidated(id, principal)
}

// satisfiesPrincipalValidated takes as arguments the identity and the principal.
// The function returns an error if one occurred.
// The function implements the behavior of an MSP up to and including v1.1.
func (msp *clmsp) satisfiesPrincipalValidated(id Identity, principal *m.MSPPrincipal) error {
	switch principal.PrincipalClassification {
	// in this case, we have to check whether the
	// identity has a role in the msp - member or admin
	case m.MSPPrincipal_ROLE:
		// Principal contains the msp role
		mspRole := &m.MSPRole{}
		err := proto.Unmarshal(principal.Principal, mspRole)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal MSPRole from principal")
		}

		// at first, we check whether the MSP
		// identifier is the same as that of the identity
		if mspRole.MspIdentifier != msp.name {
			return errors.Errorf("the identity is a member of a different MSP (expected %s, got %s)", mspRole.MspIdentifier, id.GetMSPIdentifier())
		}

		// now we validate the different msp roles
		switch mspRole.Role {
		case m.MSPRole_MEMBER:
			// in the case of member, we simply check
			// whether this identity is valid for the MSP
			mspLogger.Debugf("Checking if identity satisfies MEMBER role for %s", msp.name)
			//return msp.Validate(id)
			return nil
		case m.MSPRole_ADMIN:
			mspLogger.Debugf("Checking if identity satisfies ADMIN role for %s", msp.name)
			// in the case of admin, we check that the
			// id is exactly one of our admins
			//if id.(*clidentity).Role.Role != m.MSPRole_ADMIN {
			for _, adminidentity := range msp.admins {
				if adminidentity.Role.Role == m.MSPRole_ADMIN {
					// we do not need to check whether the admin is a valid identity
					// according to this MSP, since we already check this at Setup time
					// if there is a match, we can just return
					return nil
				}
			}
			return errors.New("This identity is not an admin")
		case m.MSPRole_CLIENT:
			fallthrough
		case m.MSPRole_PEER:
			mspLogger.Debugf("Checking if identity satisfies role [%s] for %s", m.MSPRole_MSPRoleType_name[int32(mspRole.Role)], msp.name)
			/*
				if err := msp.hasOURole(id, mspRole.Role); err != nil {
					return errors.Wrapf(err, "The identity is not a [%s] under this MSP [%s]", m.MSPRole_MSPRoleType_name[int32(mspRole.Role)], msp.name)
				}
			*/
			return nil
		default:
			return errors.Errorf("invalid MSP role type %d", int32(mspRole.Role))
		}
	case m.MSPPrincipal_IDENTITY:
		// in this case we have to deserialize the principal's identity
		// and compare it byte-by-byte with our cert
		principalId, err := msp.DeserializeIdentity(principal.Principal)
		if err != nil {
			return errors.WithMessage(err, "invalid principal, expect a serialized ibpcla identity")
		}

		if bytes.Equal(id.(*clidentity).PA, principalId.(*clidentity).PA) {
			return principalId.Validate()
		}

		return errors.New("The identities do not match")
	case m.MSPPrincipal_ORGANIZATION_UNIT:
		// Principal contains the OrganizationUnit
		OU := &m.OrganizationUnit{}
		err := proto.Unmarshal(principal.Principal, OU)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal OrganizationUnit from principal")
		}

		// at first, we check whether the MSP
		// identifier is the same as that of the identity
		if OU.MspIdentifier != msp.name {
			return errors.Errorf("the identity is a member of a different MSP (expected %s, got %s)", OU.MspIdentifier, id.GetMSPIdentifier())
		}

		// we then check if the identity is valid with this MSP
		// and fail if it is not
		if OU.OrganizationalUnitIdentifier != id.(*clidentity).OU.OrganizationalUnitIdentifier {
			return errors.Errorf("user is not part of the desired organizationalunit")
		}
		if !bytes.Equal(OU.CertifiersIdentifier, OU.CertifiersIdentifier) {
			return errors.Errorf("OU CertifiersIdentifier not match")
		}

		return nil
	case m.MSPPrincipal_COMBINED:
		if msp.version <= MSPv1_1 {
			return errors.Errorf("Combined MSP Principals are unsupported before MSPv1_2")
		}

		// Principal is a combination of multiple principals.
		principals := &m.CombinedPrincipal{}
		err := proto.Unmarshal(principal.Principal, principals)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal CombinedPrincipal from principal")
		}
		// Return an error if there are no principals in the combined principal.
		if len(principals.Principals) == 0 {
			return errors.New("no principals in CombinedPrincipal")
		}
		// Recursively call msp.SatisfiesPrincipal for all combined principals.
		// There is no limit for the levels of nesting for the combined principals.
		for _, cp := range principals.Principals {
			err = msp.satisfiesPrincipalValidated(id, cp)
			if err != nil {
				return err
			}
		}
		// The identity satisfies all the principals
		return nil
	case m.MSPPrincipal_ANONYMITY:
		if msp.version <= MSPv1_1 {
			return errors.Errorf("Anonymity MSP Principals are unsupported before MSPv1_2")
		}
		anon := &m.MSPIdentityAnonymity{}
		err := proto.Unmarshal(principal.Principal, anon)
		if err != nil {
			return errors.Wrap(err, "could not unmarshal MSPIdentityAnonymity from principal")
		}
		switch anon.AnonymityType {
		case m.MSPIdentityAnonymity_ANONYMOUS:
			return errors.New("Principal is anonymous, but CL MSP does not support anonymous identities")
		case m.MSPIdentityAnonymity_NOMINAL:
			return nil
		default:
			return errors.Errorf("Unknown principal anonymity type: %d", anon.AnonymityType)
		}
		// if no match was found , return an error
		return errors.New("The clidentities do not match")
	default:
		return errors.Errorf("invalid principal type %d", int32(principal.PrincipalClassification))
	}
}

/*
// getCertificationChain returns the certification chain of the passed identity within this msp
func (msp *clmsp) getCertificationChain(id Identity) ([]*x509.Certificate, error) {
	mspLogger.Debugf("MSP %s getting certification chain", msp.name)

	switch id := id.(type) {
	// If this identity is of this specific type,
	// this is how I can validate it given the
	// root of trust this MSP has
	case *identity:
		return msp.getCertificationChainForBCCSPIdentity(id)
	default:
		return nil, errors.New("identity type not recognized")
	}
}
*/

/*
// getCertificationChainForBCCSPIdentity returns the certification chain of the passed bccsp identity within this msp
func (msp *clmsp) getCertificationChainForBCCSPIdentity(id *identity) ([]*x509.Certificate, error) {
	if id == nil {
		return nil, errors.New("Invalid bccsp identity. Must be different from nil.")
	}

	// we expect to have a valid VerifyOptions instance
	if msp.opts == nil {
		return nil, errors.New("Invalid msp instance")
	}

	// CAs cannot be directly used as identities..
	if id.cert.IsCA {
		return nil, errors.New("An X509 certificate with Basic Constraint: " +
			"Certificate Authority equals true cannot be used as an identity")
	}

	return msp.getValidationChain(id.cert, false)
}
*/

func (msp *clmsp) getUniqueValidationChain(cert *x509.Certificate, opts x509.VerifyOptions) ([]*x509.Certificate, error) {
	// ask golang to validate the cert for us based on the options that we've built at setup time
	//skip, because no root ca exits
	/*
		if msp.opts == nil {
			return nil, errors.New("the supplied identity has no verify options")
		}
	*/

	validationChains, err := cert.Verify(opts)
	if err != nil {
		return nil, errors.WithMessage(err, "the supplied identity is not valid")
	}

	// we only support a single validation chain;
	// if there's more than one then there might
	// be unclarity about who owns the identity
	if len(validationChains) != 1 {
		return nil, errors.Errorf("this MSP only supports a single validation chain, got %d", len(validationChains))
	}

	return validationChains[0], nil
}

/*
func (msp *clmsp) getValidationChain(cert *x509.Certificate, isIntermediateChain bool) ([]*x509.Certificate, error) {
	validationChain, err := msp.getUniqueValidationChain(cert, msp.getValidityOptsForCert(cert))
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting validation chain")
	}

	// we expect a chain of length at least 2
	if len(validationChain) < 2 {
		return nil, errors.Errorf("expected a chain of length at least 2, got %d", len(validationChain))
	}

	// check that the parent is a leaf of the certification tree
	// if validating an intermediate chain, the first certificate will the parent
	parentPosition := 1
	if isIntermediateChain {
		parentPosition = 0
	}
	if msp.certificationTreeInternalNodesMap[string(validationChain[parentPosition].Raw)] {
		return nil, errors.Errorf("invalid validation chain. Parent certificate should be a leaf of the certification tree [%v]", cert.Raw)
	}
	return validationChain, nil
}
*/

// getPAIdentifier returns the certification chain identifier of the passed identity within this msp.
// The identifier is computes as the SHA256 of the concatenation of the certificates in the chain.
func (msp *clmsp) getPAIdentifier(id Identity) ([]byte, error) {
	scid := id.(*clidentity).id.Id
	if len(scid) == 0 {
		return nil, errors.New(fmt.Sprintf("failed getting PA identifier for [%v]", id))
	}

	cid, err := hex.DecodeString(scid)
	return cid, err
}

/*
func (msp *clmsp) getKGCIdentifierFromChain(chain []*x509.Certificate) ([]byte, error) {
	// Hash the chain
	// Use the hash of the identity's certificate as id in the IdentityIdentifier
	hashOpt, err := bccsp.GetHashOpt(msp.cryptoConfig.IdentityIdentifierHashFunction)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function options")
	}

	hf, err := msp.csp.GetHash(hashOpt)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function when computing certification chain identifier")
	}
	for i := 0; i < len(chain); i++ {
		hf.Write(chain[i].Raw)
	}
	return hf.Sum(nil), nil
}
*/

// sanitizeCert ensures that x509 certificates signed using ECDSA
// do have signatures in Low-S. If this is not the case, the certificate
// is regenerated to have a Low-S signature.
func (msp *clmsp) sanitizeCert(cert *x509.Certificate) (*x509.Certificate, error) {
	if isECDSASignedCert(cert) {
		// Lookup for a parent certificate to perform the sanitization
		var parentCert *x509.Certificate
		chain, err := msp.getUniqueValidationChain(cert, msp.getValidityOptsForCert(cert))
		if err != nil {
			return nil, err
		}

		// at this point, cert might be a root CA certificate
		// or an intermediate CA certificate
		if cert.IsCA && len(chain) == 1 {
			// cert is a root CA certificate
			parentCert = cert
		} else {
			parentCert = chain[1]
		}

		// Sanitize
		cert, err = sanitizeECDSASignedCert(cert, parentCert)
		if err != nil {
			return nil, err
		}
	}
	return cert, nil
}

// IsWellFormed checks if the given identity can be deserialized into its provider-specific form.
// In this MSP implementation, well formed means that it contains a
// marshaled SerializedIdemixIdentity protobuf message.
func (msp *clmsp) IsWellFormed(identity *m.SerializedIdentity) error {
	sId := new(m.SerializedIBPCLAIdentity)
	err := proto.Unmarshal(identity.IdBytes, sId)
	if err != nil {
		return errors.Wrap(err, "not an IBPCLA identity")
	}
	return nil
}

// ADD BY WYH
// GetSigningIdentity returns a specific signing
// identity identified by the supplied identifier
func (msp *clmsp) GetBccsp(identifier string) (bccsp.BCCSP, error) {
	// TODO
	if msp.csp != nil {
		return msp.csp, nil
	}

	return nil, fmt.Errorf("No bccsp identity for %s", identifier)

}

func (msp *clmsp) getValidityOptsForCert(cert *x509.Certificate) x509.VerifyOptions {
	// First copy the opts to override the CurrentTime field
	// in order to make the certificate passing the expiration test
	// independently from the real local current time.
	// This is a temporary workaround for FAB-3678

	var tempOpts x509.VerifyOptions
	tempOpts.Roots = msp.opts.Roots
	tempOpts.DNSName = msp.opts.DNSName
	tempOpts.Intermediates = msp.opts.Intermediates
	tempOpts.KeyUsages = msp.opts.KeyUsages
	tempOpts.CurrentTime = cert.NotBefore.Add(time.Second)

	return tempOpts
}
