/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"encoding/pem"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/protos/msp"
	m "github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

var mspclIdentityLogger = flogging.MustGetLogger("msp.clidentity")

type clidentity struct {

	// the vice identity of this identity
	PA []byte

	// id contains the identifier (MSPID and identity identifier) for this instance
	id *IdentityIdentifier

	// reference to the MSP that "owns" this identity
	msp *clmsp
}

func newclIdentity(PA []byte, msp *clmsp) (Identity, error) {
	if mspclIdentityLogger.IsEnabledFor(zapcore.DebugLevel) {
		mspclIdentityLogger.Debugf("Creating identity instance for KGC")
	}

	if PA == nil {
		return nil, errors.New("failed getting PA")
	}

	// Use the hash of the identity's certificate as id in the IdentityIdentifier
	hashOpt, err := bccsp.GetHashOpt(msp.cryptoConfig.IdentityIdentifierHashFunction)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function options")
	}

	// Compute identity identifier
	digest, err := msp.csp.Hash(PA, hashOpt)
	if err != nil {
		return nil, errors.WithMessage(err, "failed hashing raw pubs to compute the id of the IdentityIdentifier")
	}

	id := &IdentityIdentifier{
		Mspid: msp.name,
		Id:    hex.EncodeToString(digest)}

	return &clidentity{PA: PA, id: id, msp: msp}, nil
}

// ExpiresAt returns the time at which the Identity expires.
func (id *clidentity) ExpiresAt() time.Time {
	//IBPCLA MSP currently does not use expiration dates, so return zero time
	return time.Time{}
}

// SatisfiesPrincipal returns null if this instance matches the supplied principal or an error otherwise
func (id *clidentity) SatisfiesPrincipal(principal *msp.MSPPrincipal) error {
	return id.msp.SatisfiesPrincipal(id, principal)
}

// GetIdentifier returns the identifier (MSPID/IDID) for this instance
func (id *clidentity) GetIdentifier() *IdentityIdentifier {
	return id.id
}

// GetMSPIdentifier returns the MSP identifier for this instance
func (id *clidentity) GetMSPIdentifier() string {
	return id.id.Mspid
}

// Validate returns nil if this instance is a valid clidentity or an error otherwise
func (id *clidentity) Validate() error {
	return id.msp.Validate(id)
}

// GetOrganizationalUnits returns the OU for this instance
func (id *clidentity) GetOrganizationalUnits() []*OUIdentifier {

	if id.msp.ouIdentifiers == nil {
		mspIdentityLogger.Errorf("Failed to get OrganizationalUnitIdentifier in GetOrganizationalUnits")
		return nil
	}

	res := []*OUIdentifier{}

	for k, v := range id.msp.ouIdentifiers {
		if v == nil {
			mspIdentityLogger.Errorf("Failed to get OrganizationalUnitIdentifier in GetOrganizationalUnits")
			return nil
		}
		res = append(res, &OUIdentifier{
			CertifiersIdentifier:         v[0],
			OrganizationalUnitIdentifier: k,
		})

		cid, err := id.msp.getPAIdentifier(id)
		if err != nil {
			mspIdentityLogger.Errorf("Failed to get PA Identifier")
			return nil
		}
		res = append(res, &OUIdentifier{
			CertifiersIdentifier:         cid,
			OrganizationalUnitIdentifier: k,
		})
	}

	if len(res) == 0 {
		return nil
	}

	return res
}

// Anonymous returns true if this clidentity provides anonymity
func (id *clidentity) Anonymous() bool {
	return false
}

// NewSerializedIdentity returns a serialized clidentity
// having as content the passed mspID and x509 certificate in PEM format.
// This method does not check the validity of certificate nor
// any consistency of the mspID with it.
func NewSerializedclIdentity(mspID string, certPEM []byte) ([]byte, error) {
	// We serialize identities by prepending the MSPID
	// and appending the x509 cert in PEM format
	bl, _ := pem.Decode(certPEM)
	serialized := &m.SerializedIBPCLAIdentity{}
	serialized.PA = bl.Bytes

	CLIDBytes, err := proto.Marshal(serialized)

	sId := &msp.SerializedIdentity{Mspid: mspID, IdBytes: CLIDBytes}
	raw, err := proto.Marshal(sId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed serializing clidentity [%s][%X]", mspID, certPEM)
	}
	return raw, nil
}

// Verify checks against a signature and a message
// to determine whether this clidentity produced the
// signature; it returns nil if so or an error otherwise
func (id *clidentity) Verify(msg []byte, sig []byte) error {
	//mspclIdentityLogger.Infof("Verifying signature")

	// Compute Hash
	hashOpt, err := id.getHashOpt(id.msp.cryptoConfig.SignatureHashFamily)
	if err != nil {
		return errors.WithMessage(err, "failed getting hash function options")
	}

	digest, err := id.msp.csp.Hash(msg, hashOpt)
	if err != nil {
		return errors.WithMessage(err, "failed computing digest")
	}

	//hard-coding to SHA2 according to clgen
	hashOptID, err := id.getHashOpt(bccsp.SHA2)

	//compute HID
	var buffer bytes.Buffer
	buffer.Write([]byte(id.id.Mspid))
	buffer.Write(id.PA)
	HID, err := id.msp.csp.Hash(buffer.Bytes(), hashOptID)
	if err != nil {
		return errors.WithMessage(err, "failed computing HID")
	}

	if mspclIdentityLogger.IsEnabledFor(zapcore.DebugLevel) {
		mspclIdentityLogger.Debugf("IBPCLA Verify: digest = %s", hex.Dump(digest))
		mspclIdentityLogger.Debugf("IBPCLA Verify: sig = %s", hex.Dump(sig))
	}
	/*
		mspclIdentityLogger.Infof("+++++++++++++++++IBPCLA Verify: digest = %s", hex.Dump(digest))
		mspclIdentityLogger.Infof("+++++++++++++++++IBPCLA Verify: sig = %s", hex.Dump(sig))
		mspclIdentityLogger.Infof("+++++++++++++++++IBPCLA Verify: ID = %s", id.id.Mspid)
		mspclIdentityLogger.Infof("+++++++++++++++++IBPCLA Verify: HID = %s", hex.Dump(HID))
		mspclIdentityLogger.Infof("+++++++++++++++++IBPCLA Verify: PA = %s", hex.Dump(id.PA))
	*/
	//verify signature
	valid, err := id.msp.csp.Verify(
		nil,
		sig,
		digest,
		&bccsp.CLVerifierOpts{
			KGCPublicKey: id.msp.rootPubs[0],
			HID:          HID,
			PA:           id.PA,
		},
	)
	if err != nil {
		return errors.WithMessage(err, "could not determine the validity of the signature")
	} else if !valid {
		return errors.New("The signature is invalid")
	}

	return nil
}

// Serialize returns a byte array representation of this clidentity
func (id *clidentity) Serialize() ([]byte, error) {
	// mspclIdentityLogger.Infof("Serializing clidentity %s", id.id)
	serialized := &m.SerializedIBPCLAIdentity{}

	serialized.PA = id.PA

	CLIDBytes, err := proto.Marshal(serialized)
	if err != nil {
		return nil, errors.Wrapf(err, "could not marshal a IdBytes for clidentity")
	}

	sId := &msp.SerializedIdentity{Mspid: id.id.Mspid, IdBytes: CLIDBytes}
	idBytes, err := proto.Marshal(sId)
	if err != nil {
		return nil, errors.Wrapf(err, "could not marshal a SerializedIdentity structure for clidentity %s", id.id)
	}

	return idBytes, nil
}

func (id *clidentity) getHashOpt(hashFamily string) (bccsp.HashOpts, error) {
	switch hashFamily {
	case bccsp.SHA2:
		return bccsp.GetHashOpt(bccsp.SHA256)
	case bccsp.SHA3:
		return bccsp.GetHashOpt(bccsp.SHA3_256)
	}
	return nil, errors.Errorf("hash familiy not recognized [%s]", hashFamily)
}

type clsigningidentity struct {
	// we embed everything from a base identity
	*clidentity

	// signer corresponds to the object that can produce signatures from this identity
	signer crypto.Signer
}

func newCLSigningIdentity(PA []byte, signer crypto.Signer, msp *clmsp) (SigningIdentity, error) {
	//mspclIdentityLogger.Infof("Creating cl signing identity instance for ID %s", id)
	block, _ := pem.Decode(PA)
	if block == nil {
		return nil, errors.New("invalid PA, failed decoding pem Bytes")

	}
	mspId, err := newclIdentity(block.Bytes, msp)
	if err != nil {
		return nil, err
	}
	return &clsigningidentity{clidentity: mspId.(*clidentity), signer: signer}, nil
}

// Sign produces a signature over msg, signed by this instance
func (id *clsigningidentity) Sign(msg []byte) ([]byte, error) {
	//mspclIdentityLogger.Infof("Signing message")

	// Compute Hash
	hashOpt, err := id.getHashOpt(id.msp.cryptoConfig.SignatureHashFamily)
	if err != nil {
		return nil, errors.WithMessage(err, "failed getting hash function options")
	}

	digest, err := id.msp.csp.Hash(msg, hashOpt)
	if err != nil {
		return nil, errors.WithMessage(err, "failed computing digest")
	}

	if len(msg) < 32 {
		mspclIdentityLogger.Debugf("Sign: plaintext: %X \n", msg)
	} else {
		mspclIdentityLogger.Debugf("Sign: plaintext: %X...%X \n", msg[0:16], msg[len(msg)-16:])
	}
	mspclIdentityLogger.Debugf("Sign: digest: %X \n", digest)

	// Sign
	return id.signer.Sign(rand.Reader, digest, nil)
}

// GetPublicVersion returns the public version of this identity,
// namely, the one that is only able to verify messages and not sign them
func (id *clsigningidentity) GetPublicVersion() Identity {
	return id.clidentity
}

func (id *clidentity) validateIdentity() error {
	// Check that the identity's OUs are compatible with those recognized by this MSP,
	// meaning that the intersection is not empty.
	if len(id.msp.ouIdentifiers) > 0 {
		found := false

		for _, OU := range id.GetOrganizationalUnits() {
			certificationIDs, exists := id.msp.ouIdentifiers[OU.OrganizationalUnitIdentifier]

			if exists {
				for _, certificationID := range certificationIDs {
					if bytes.Equal(certificationID, OU.CertifiersIdentifier) {
						found = true
						break
					}
				}
			}
		}

		if !found {
			if len(id.GetOrganizationalUnits()) == 0 {
				return errors.New("the identity certificate does not contain an Organizational Unit (OU)")
			}
			return errors.Errorf("none of the identity's organizational units [%v] are in MSP %s", id.GetOrganizationalUnits(), id.msp.name)
		}
	}

	return nil
}