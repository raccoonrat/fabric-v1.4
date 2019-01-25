/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package cryptocl

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/pkg/errors"
)

type Signer struct {
}

func (v *Signer) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	if k == nil {
		return nil, errors.New("invalid key, can not be nil")
	}

	privK, ok := k.(*signKey)
	if !ok {
		return nil, errors.New("invalid key, expected *cryptocl.signKey")
	}

	//sign
	return signECDSA(privK.privKey, digest, opts)
}

func signECDSA(k *ecdsa.PrivateKey, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
	if err != nil {
		return nil, err
	}

	s, _, err = utils.ToLowS(&k.PublicKey, s)
	if err != nil {
		return nil, err
	}

	return utils.MarshalECDSASignature(r, s)
}
