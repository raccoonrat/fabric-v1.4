/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package csp

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/signer"
	"github.com/hyperledger/fabric/common/tools/clgen/kgc"
)

// LoadPrivateKey loads a private key from file in keystorePath
func LoadPrivateKey(keystorePath string) (bccsp.Key, crypto.Signer, error) {
	var err error
	var priv bccsp.Key
	var s crypto.Signer

	opts := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystorePath,
			},
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err != nil {
		return nil, nil, err
	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "_sk") {
			rawKey, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			block, _ := pem.Decode(rawKey)
			priv, err = csp.KeyImport(block.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: true})
			if err != nil {
				return err
			}

			s, err = signer.New(csp, priv)
			if err != nil {
				return err
			}

			return nil
		}
		return nil
	}

	err = filepath.Walk(keystorePath, walkFunc)
	if err != nil {
		return nil, nil, err
	}

	return priv, s, err
}

// GeneratePrivateKey creates a private key and stores it in keystorePath
func GeneratePrivateKey(keystorePath string) (bccsp.Key,
	crypto.Signer, error) {

	var err error
	var priv bccsp.Key
	var s crypto.Signer

	opts := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystorePath,
			},
		},
	}
	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// generate a key
		priv, err = csp.KeyGen(&bccsp.ECDSAP256KeyGenOpts{Temporary: false})
		if err == nil {
			// create a crypto.Signer
			s, err = signer.New(csp, priv)
		}
	}
	return priv, s, err
}

func GetECPublicKey(priv bccsp.Key) (*ecdsa.PublicKey, error) {

	// get the public key
	pubKey, err := priv.PublicKey()
	if err != nil {
		return nil, err
	}
	// marshal to bytes
	pubKeyBytes, err := pubKey.Bytes()
	if err != nil {
		return nil, err
	}
	// unmarshal using pkix
	ecPubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}
	return ecPubKey.(*ecdsa.PublicKey), nil
}

// KGCGeneratePrivateKey creates a master key and stores it in keystorePath
func KGCGeneratePrivateKey(keystorePath string) (bccsp.Key, error) {

	var err error
	var priv bccsp.Key

	opts := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystorePath,
			},
		},
	}
	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// generate a key
		priv, err = csp.KeyGen(&bccsp.ECDSAP256KeyGenOpts{Temporary: false})
	}
	return priv, err
}

// KGCGetECPublicKey gets a master pubkey from private key and stores it in keystorePath
func KGCGetECPublicKey(priv bccsp.Key, name, keystorePath string) (*ecdsa.PublicKey, []byte, error) {

	// get the public key
	pubKey, err := priv.PublicKey()
	if err != nil {
		return nil, nil, err
	}
	// marshal to bytes
	pubKeyBytes, err := pubKey.Bytes()
	if err != nil {
		return nil, nil, err
	}

	//write pubkey out to file
	fileName := filepath.Join(keystorePath, name+"-pubkey.pem")
	pubkeyFile, err := os.Create(fileName)
	if err != nil {
		return nil, nil, err
	}
	//pem encode the cert
	err = pem.Encode(pubkeyFile, &pem.Block{Type: "ECC PUBLIC KEY", Bytes: pubKeyBytes})
	certFile.Close()
	if err != nil {
		return nil, nil, err
	}
	// unmarshal using pkix
	ecPubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, nil, err
	}
	return ecPubKey.(*ecdsa.PublicKey), pubKeyBytes, nil
}

func GenFinalKeyPair(keystorePath, name string, x *ecdsa.PrivateKey, partialkey *kgc.PartialKey) error {

	var finalPrivateKey *ecdsa.PrivateKey
	finalPrivateKey, err = GenFinalKeyPairInternal(name, x, partialkey)
	if err != nil {
		return err
	}

	//write Final private key to file
	fileName := filepath.Join(keystorePath, ID+"-PA.pem")
	PAFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	//pem encode the public key
	err = pem.Encode(PAFile, &pem.Block{Type: "ECC PUBLIC KEY", Bytes: partialkey.PartialPublickKey.Bytes()})
	PAFile.Close()
	if err != nil {
		return nil, err
	}
	return nil
}

func GenFinalKeyPairInternal(ID string, ClientPrivateKey *ecdsa.PrivateKey, partialkey *kgc.PartialKey) (*big.Int, error) {

	var buffer bytes.Buffer
	n := ClientPrivateKey.Curve.Params().N
	//e=h(ID||PA)
	buffer.Write([]byte(ID))
	buffer.Write(PA.Bytes())
	e := sha256.Sum256(buffer)
	e0 := new(big.Int).SetBytes(e[0:15])

	//d=e0*x+za
	d = new(big.Int).Add()
	return
}
