/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package cryptocl

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/pkg/errors"
)

type KGCKeyImporter struct {
}

func (v *KGCKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	pembyte, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("invalid raw, expected byte array")
	}

	if len(pembyte) == 0 {
		return nil, errors.New("invalid raw, it must not be nil")
	}

	block, _ := pem.Decode(pembyte)
	if block == nil {
		return nil, errors.New("invalid key, pem decode failed")
	}

	ecPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Invalid raw material. Expected *Block")
	}

	pk, ok := ecPubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Invalid pub key material. Expected *ecdsa.PublicKey.")
	}

	return &rootPublicKey{pk}, nil

}

type SignerKeyImporter struct {
}

func (v *SignerKeyImporter) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	der, ok := raw.([]byte)
	if !ok {
		return nil, errors.New("[ECDSADERPrivateKeyImportOpts] Invalid raw material. Expected byte array.")
	}

	if len(der) == 0 {
		return nil, errors.New("[ECDSADERPrivateKeyImportOpts] Invalid raw. It must not be nil.")
	}

	lowLevelKey, err := utils.DERToPrivateKey(der)
	if err != nil {
		return nil, fmt.Errorf("Failed converting PKIX to ECDSA public key [%s]", err)
	}

	ecdsaSK, ok := lowLevelKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("Failed casting to ECDSA private key. Invalid raw material.")
	}

	return &signKey{ecdsaSK}, nil
}
