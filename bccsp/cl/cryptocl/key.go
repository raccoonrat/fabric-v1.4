/*
SPDX-License-Identifier: Apache-2.0
*/
package cryptocl

import (
	"crypto/ecdsa"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
)

// rootPublicKey contains the kgc root public key
// and implements the bccsp.Key interface
type rootPublicKey struct {
	pubKey *ecdsa.PublicKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k rootPublicKey) Bytes() (raw []byte, err error) {
	raw, err = x509.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k rootPublicKey) SKI() []byte {
	return nil
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k rootPublicKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k rootPublicKey) Private() bool {
	return false
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k rootPublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

// signKey contains the private key of the signing identity
// and implements the bccsp.Key interface
type signKey struct {
	privKey *ecdsa.PrivateKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k signKey) Bytes() (raw []byte, err error) {
	return nil, errors.New("Not supported.")
}

// SKI returns the subject key identifier of this key.
func (k signKey) SKI() []byte {
	return nil
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k signKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k signKey) Private() bool {
	return true
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k signKey) PublicKey() (bccsp.Key, error) {
	return &rootPublicKey{&k.privKey.PublicKey}, nil
}

// NewCecretKey returns a new signKey struct
func NewSecretKey() *signKey {
	return &signKey{}
}
