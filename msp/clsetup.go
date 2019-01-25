/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/bccsp"
	m "github.com/hyperledger/fabric/protos/msp"
	errors "github.com/pkg/errors"
)

func (msp *clmsp) getKGCIdentifier(PubRaw []byte) ([]byte, error) {

	//decode and get pub
	/*
		bl, _ := pem.Decode(PubRaw)
		if bl == nil {
			return nil, errors.New("invalid KGC Pubs, pem decode fail")
		}
	*/
	found := false
	// Search among root certificates
	var temp []byte
	for _, v := range msp.rootPubs {
		temp, _ = v.Bytes()
		if bytes.Equal(temp, PubRaw) {
			found = true
			break
		}
	}
	if !found {
		// kgc Pub not valid, reject configuration
		return nil, fmt.Errorf("Failed adding OU. Pub [%v] not in root pubs.", PubRaw)
	}

	// compute the hash of the pub
	hashOpt, err := bccsp.GetHashOpt(msp.cryptoConfig.IdentityIdentifierHashFunction)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function options    ")
	}
	hf, err := msp.csp.GetHash(hashOpt)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function when computing kgc identifier")
	}
	hf.Write(PubRaw)

	return hf.Sum(nil), nil

}

func (msp *clmsp) setupCrypto(conf *m.CLMSPConfig) error {
	msp.cryptoConfig = conf.CryptoConfig
	if msp.cryptoConfig == nil {
		// Move to defaults
		msp.cryptoConfig = &m.FabricCryptoConfig{
			SignatureHashFamily:            bccsp.SHA2,
			IdentityIdentifierHashFunction: bccsp.SHA256,
		}
		mspLogger.Debugf("CryptoConfig was nil. Move to defaults.")
	}
	if msp.cryptoConfig.SignatureHashFamily == "" {
		msp.cryptoConfig.SignatureHashFamily = bccsp.SHA2
		mspLogger.Debugf("CryptoConfig.SignatureHashFamily was nil. Move to defaults.")
	}
	if msp.cryptoConfig.IdentityIdentifierHashFunction == "" {
		msp.cryptoConfig.IdentityIdentifierHashFunction = bccsp.SHA256
		mspLogger.Debugf("CryptoConfig.IdentityIdentifierHashFunction was nil. Move to defaults.")
	}

	return nil
}

func (msp *clmsp) setupKGCs(conf *m.CLMSPConfig) error {
	// make and fill the set of KGC pubs - we expect them to be there
	if len(conf.KGCPubs) == 0 {
		return errors.New("expected at least one KGC pubs")
	}

	// make and fill the set of KGC Public keys
	msp.rootPubs = make([]bccsp.Key, len(conf.KGCPubs))

	for i, v := range conf.KGCPubs {
		// get the KGC public key in the right format
		PubK, err := msp.csp.KeyImport(v, &bccsp.CLKGCPublicKeyImportOpts{Temporary: true})
		if err != nil {
			return errors.WithMessage(err, "failed to import KGC Public key")
		}
		msp.rootPubs[i] = PubK
	}

	return nil
}

func (msp *clmsp) setupAdmins(conf *m.CLMSPConfig) error {
	// make and fill the set of admin PAs (if present)
	msp.admins = make([]clidentity, len(conf.Admins))
	for i, admconf := range conf.Admins {
		id, err := msp.getclAdminIdentityFromConf(admconf)
		if err != nil {
			return err
		}
		msp.admins[i] = *id
	}

	return nil
}

func (msp *clmsp) setupCRLs(conf *m.CLMSPConfig) error {
	// setup the CRL (if present)
	msp.CRL = make([]*pkix.CertificateList, len(conf.RevocationList))
	for i, crlbytes := range conf.RevocationList {
		crl, err := x509.ParseCRL(crlbytes)
		if err != nil {
			return errors.Wrap(err, "could not parse RevocationList")
		}

		// TODO: pre-verify the signature on the CRL and create a map
		//       of CA certs to respective CRLs so that later upon
		//       validation we can already look up the CRL given the
		//       chain of the certificate to be validated

		msp.CRL[i] = crl
	}

	return nil
}

func (msp *clmsp) setupNodeOUs(config *m.CLMSPConfig) error {
	if config.FabricNodeOus != nil {

		msp.ouEnforcement = config.FabricNodeOus.Enable

		// ClientOU
		msp.clientOU = &OUIdentifier{OrganizationalUnitIdentifier: config.FabricNodeOus.ClientOuIdentifier.OrganizationalUnitIdentifier}
		if len(config.FabricNodeOus.ClientOuIdentifier.Certificate) != 0 {
			certifiersIdentifier, err := msp.getKGCIdentifier(config.FabricNodeOus.ClientOuIdentifier.Certificate)
			if err != nil {
				return err
			}
			msp.clientOU.CertifiersIdentifier = certifiersIdentifier
		}

		// PeerOU
		msp.peerOU = &OUIdentifier{OrganizationalUnitIdentifier: config.FabricNodeOus.PeerOuIdentifier.OrganizationalUnitIdentifier}
		if len(config.FabricNodeOus.PeerOuIdentifier.Certificate) != 0 {
			certifiersIdentifier, err := msp.getKGCIdentifier(config.FabricNodeOus.PeerOuIdentifier.Certificate)
			if err != nil {
				return err
			}
			msp.peerOU.CertifiersIdentifier = certifiersIdentifier
		}

	} else {
		msp.ouEnforcement = false
	}

	return nil
}

func (msp *clmsp) setupSigningIdentity(conf *m.CLMSPConfig) error {
	if conf.CLSigningIdentity != nil {
		sid, err := msp.getSigningIdentityFromConf(conf.CLSigningIdentity)
		if err != nil {
			return err
		}
		msp.signer = sid
	}

	return nil
}

func (msp *clmsp) setupOUs(conf *m.CLMSPConfig) error {

	fmt.Println("---setup OUs")
	msp.ouIdentifiers = make(map[string][][]byte)
	for _, ou := range conf.OrganizationalUnitIdentifiers {

		kgcIdentifier, err := msp.getKGCIdentifier(ou.Certificate)
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("failed getting kgcpub for [%v]", ou))
		}

		// Check for duplicates
		found := false
		for _, id := range msp.ouIdentifiers[ou.OrganizationalUnitIdentifier] {
			if bytes.Equal(id, kgcIdentifier) {
				mspLogger.Warningf("Duplicate found in ou identifiers [%s, %v]", ou.OrganizationalUnitIdentifier, id)
				found = true
				break
			}
		}

		if !found {
			// No duplicates found, add it
			msp.ouIdentifiers[ou.OrganizationalUnitIdentifier] = append(
				msp.ouIdentifiers[ou.OrganizationalUnitIdentifier],
				kgcIdentifier,
			)
		}
	}

	return nil
}

func (msp *clmsp) setupTLSCAs(conf *m.CLMSPConfig) error {

	opts := &x509.VerifyOptions{Roots: x509.NewCertPool(), Intermediates: x509.NewCertPool()}

	// Load TLS root and intermediate CA identities
	msp.tlsRootCerts = make([][]byte, len(conf.TlsRootCerts))
	rootCerts := make([]*x509.Certificate, len(conf.TlsRootCerts))
	for i, trustedCert := range conf.TlsRootCerts {
		cert, err := msp.getCertFromPem(trustedCert)
		if err != nil {
			return err
		}

		rootCerts[i] = cert
		msp.tlsRootCerts[i] = trustedCert
		opts.Roots.AddCert(cert)
	}

	// make and fill the set of intermediate certs (if present)
	msp.tlsIntermediateCerts = make([][]byte, len(conf.TlsIntermediateCerts))
	intermediateCerts := make([]*x509.Certificate, len(conf.TlsIntermediateCerts))
	for i, trustedCert := range conf.TlsIntermediateCerts {
		cert, err := msp.getCertFromPem(trustedCert)
		if err != nil {
			return err
		}

		intermediateCerts[i] = cert
		msp.tlsIntermediateCerts[i] = trustedCert
		opts.Intermediates.AddCert(cert)
	}

	// ensure that our CAs are properly formed and that they are valid
	for _, cert := range append(append([]*x509.Certificate{}, rootCerts...), intermediateCerts...) {
		if cert == nil {
			continue
		}

		if !cert.IsCA {
			return errors.Errorf("CA Certificate did not have the CA attribute, (SN: %x)", cert.SerialNumber)
		}
		if _, err := getSubjectKeyIdentifierFromCert(cert); err != nil {
			return errors.WithMessage(err, fmt.Sprintf("CA Certificate problem with Subject Key Identifier extension, (SN: %x)", cert.SerialNumber))
		}

		if err := msp.validateTLSCAIdentity(cert, opts); err != nil {
			return errors.WithMessage(err, fmt.Sprintf("CA Certificate is not valid, (SN: %s)", cert.SerialNumber))
		}
	}

	return nil
}

func (msp *clmsp) setupV1(conf1 *m.CLMSPConfig) error {
	err := msp.preSetupV1(conf1)
	if err != nil {
		return err
	}

	err = msp.postSetupV1(conf1)
	if err != nil {
		return err
	}

	return nil
}

func (msp *clmsp) preSetupV1(conf *m.CLMSPConfig) error {
	// setup crypto config
	if err := msp.setupCrypto(conf); err != nil {
		return errors.Wrap(err, "setup crypto error")
	}

	// Setup KGCs
	if err := msp.setupKGCs(conf); err != nil {
		return errors.Wrap(err, "setup KGC error")

	}

	// Setup Admins
	if err := msp.setupAdmins(conf); err != nil {
		return errors.Wrap(err, "setup Admins error")
	}

	// Setup CRLs
	if err := msp.setupCRLs(conf); err != nil {
		return errors.Wrap(err, "setup CRL error")
	}

	// setup the signer (if present)
	if err := msp.setupSigningIdentity(conf); err != nil {
		return errors.Wrap(err, "setup signing identity error")
	}

	// setup TLS CAs
	if err := msp.setupTLSCAs(conf); err != nil {
		return err
	}

	/*
		// setup the OUs
		if err := msp.setupOUs(conf); err != nil {
			return err
		}
	*/
	return nil
}

func (msp *clmsp) postSetupV1(conf *m.CLMSPConfig) error {
	// make sure that admins are valid members as well
	// this way, when we validate an admin MSP principal
	// we can simply check for exact match of certs
	for i, admin := range msp.admins {
		err := admin.Validate()
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("admin %d is invalid", i))
		}
	}

	return nil
}

func (msp *clmsp) setupV11(conf *m.CLMSPConfig) error {
	err := msp.preSetupV1(conf)
	if err != nil {
		return err
	}

	// setup NodeOUs
	if err := msp.setupNodeOUs(conf); err != nil {
		return err
	}

	err = msp.postSetupV11(conf)
	if err != nil {
		return err
	}

	return nil
}

func (msp *clmsp) postSetupV11(conf *m.CLMSPConfig) error {
	// Check for OU enforcement
	if !msp.ouEnforcement {
		// No enforcement required. Call post setup as per V1
		return msp.postSetupV1(conf)
	}

	// Check that admins are clients
	principalBytes, err := proto.Marshal(&m.MSPRole{Role: m.MSPRole_CLIENT, MspIdentifier: msp.name})
	if err != nil {
		return errors.Wrapf(err, "failed creating MSPRole_CLIENT")
	}
	principal := &m.MSPPrincipal{
		PrincipalClassification: m.MSPPrincipal_ROLE,
		Principal:               principalBytes}
	for i, admin := range msp.admins {
		err = admin.SatisfiesPrincipal(principal)
		if err != nil {
			return errors.WithMessage(err, fmt.Sprintf("admin %d is invalid", i))
		}
	}

	return nil
}

func (msp *clmsp) validateTLSCAIdentity(cert *x509.Certificate, opts *x509.VerifyOptions) error {
	if !cert.IsCA {
		return errors.New("Only CA identities can be validated")
	}

	validationChain, err := msp.getUniqueValidationChain(cert, *opts)
	if err != nil {
		return errors.WithMessage(err, "could not obtain certification chain")
	}
	if len(validationChain) == 1 {
		// validationChain[0] is the root CA certificate
		return nil
	}

	return msp.validateCertAgainstChain(cert, validationChain)
}

func (msp *clmsp) validateCertAgainstChain(cert *x509.Certificate, validationChain []*x509.Certificate) error {
	// here we know that the identity is valid; now we have to check whether it has been revoked

	// identify the SKI of the CA that signed this cert
	SKI, err := getSubjectKeyIdentifierFromCert(validationChain[1])
	if err != nil {
		return errors.WithMessage(err, "could not obtain Subject Key Identifier for signer cert")
	}

	// check whether one of the CRLs we have has this cert's
	// SKI as its AuthorityKeyIdentifier
	for _, crl := range msp.CRL {
		aki, err := getAuthorityKeyIdentifierFromCrl(crl)
		if err != nil {
			return errors.WithMessage(err, "could not obtain Authority Key Identifier for crl")
		}

		// check if the SKI of the cert that signed us matches the AKI of any of the CRLs
		if bytes.Equal(aki, SKI) {
			// we have a CRL, check whether the serial number is revoked
			for _, rc := range crl.TBSCertList.RevokedCertificates {
				if rc.SerialNumber.Cmp(cert.SerialNumber) == 0 {
					// We have found a CRL whose AKI matches the SKI of
					// the CA (root or intermediate) that signed the
					// certificate that is under validation. As a
					// precaution, we verify that said CA is also the
					// signer of this CRL.
					err = validationChain[1].CheckCRLSignature(crl)
					if err != nil {
						// the CA cert that signed the certificate
						// that is under validation did not sign the
						// candidate CRL - skip
						mspLogger.Warningf("Invalid signature over the identified CRL, error %+v", err)
						continue
					}

					// A CRL also includes a time of revocation so that
					// the CA can say "this cert is to be revoked starting
					// from this time"; however here we just assume that
					// revocation applies instantaneously from the time
					// the MSP config is committed and used so we will not
					// make use of that field
					return errors.New("The certificate has been revoked")
				}
			}
		}
	}

	return nil
}
