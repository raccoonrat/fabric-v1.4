/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package csp

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/signer"
	"github.com/pkg/errors"
)

type pkcs8Info struct {
	Version             int
	PrivateKeyAlgorithm []asn1.ObjectIdentifier
	PrivateKey          []byte
}

type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

var (
	oidNamedCurveP224 = asn1.ObjectIdentifier{1, 3, 132, 0, 33}
	oidNamedCurveP256 = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	oidNamedCurveP384 = asn1.ObjectIdentifier{1, 3, 132, 0, 34}
	oidNamedCurveP521 = asn1.ObjectIdentifier{1, 3, 132, 0, 35}
)

var oidPublicKeyECDSA = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}

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

// LoadKGCMasterKey loads a master key from file in keystorePath
func LoadKGCMasterKey(keystorePath string) (*ecdsa.PrivateKey, error) {
	var err error
	var priv *ecdsa.PrivateKey

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
		return nil, err
	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "_sk") {
			rawKey, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			block, _ := pem.Decode(rawKey)
			bpriv, err := csp.KeyImport(block.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: true})
			if err != nil {
				return err
			}

			priv, err = BccspKey2ecdsaKey(bpriv)
			if err != nil {
				return err
			}

			return nil
		}
		return nil
	}

	err = filepath.Walk(keystorePath, walkFunc)
	if (err != nil) || (priv == nil) {
		return nil, errors.Wrapf(err, "could not load a valid sk from directory %s", keystorePath)
	}

	return priv, err
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

// KGCGenerateMasterKey creates a master key and stores it in keystorePath
func KGCGenerateMasterKey(keystorePath string) (*ecdsa.PrivateKey, []byte, error) {

	//todo: support sm
	curve := elliptic.P256()
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(privKey.Public())
	if err != nil {
		return nil, nil, err
	}
	//pemBytes := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubKeyBytes})
	err = ioutil.WriteFile(filepath.Join(keystorePath, "KGC-PublicKey"), pubKeyBytes, 0644)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to store KGC public key")
	}
	err = ioutil.WriteFile(filepath.Join(keystorePath, "KGC-MasterKey"), privKey.D.Bytes(), 0644)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to store KGC master key")
	}

	return privKey, pubKeyBytes, nil
}

func BccspKey2ecdsaKey(bkey bccsp.Key) (*ecdsa.PrivateKey, error) {
	var key *ecdsa.PrivateKey
	key.PublicKey.Curve = elliptic.P256()
	buffer, err := bkey.Bytes()
	if err != nil {
		fmt.Println("error in copying buffer")
		return nil, err
	}
	key.D.SetBytes(buffer)
	return key, err
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
	err = pem.Encode(pubkeyFile, &pem.Block{Type: "PUBLIC KEY", Bytes: pubKeyBytes})
	pubkeyFile.Close()
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

func GenFinalKeyPair(ID, OU, Role string, ClientPrivateKey *ecdsa.PrivateKey, PartialPublicKey []byte, PartialPrivateKey []byte) ([]byte, error) {

	var buffer bytes.Buffer
	n := ClientPrivateKey.Params().N

	//e=h(ID||PA)
	buffer.Write([]byte(ID))
	buffer.Write(PartialPublicKey)
	buffer.Write([]byte(OU))
	buffer.Write([]byte(Role))
	e := sha256.Sum256(buffer.Bytes())
	e0 := new(big.Int).SetBytes(e[0:15])

	//d=e0*x+za
	e0.Mul(e0, ClientPrivateKey.D)
	e0.Mod(e0, n)

	za := new(big.Int).SetBytes(PartialPrivateKey)
	d := new(big.Int).Add(e0, za)
	d.Mod(d, n)

	return d.Bytes(), nil
}

func PrivateKeyToPEM(k *ecdsa.PrivateKey) ([]byte, error) {
	// get the oid for the curve
	oidNamedCurve, ok := oidFromNamedCurve(k.Curve)
	if !ok {
		return nil, errors.New("unknown elliptic curve")
	}

	// based on https://golang.org/src/crypto/x509/sec1.go
	privateKeyBytes := k.D.Bytes()
	paddedPrivateKey := make([]byte, (k.Curve.Params().N.BitLen()+7)/8)
	copy(paddedPrivateKey[len(paddedPrivateKey)-len(privateKeyBytes):], privateKeyBytes)
	// omit NamedCurveOID for compatibility as it's optional
	asn1Bytes, err := asn1.Marshal(ecPrivateKey{
		Version:    1,
		PrivateKey: paddedPrivateKey,
		PublicKey:  asn1.BitString{Bytes: elliptic.Marshal(k.Curve, k.X, k.Y)},
	})

	if err != nil {
		return nil, fmt.Errorf("error marshaling EC key to asn1 [%s]", err)
	}

	var pkcs8Key pkcs8Info
	pkcs8Key.Version = 0
	pkcs8Key.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 2)
	pkcs8Key.PrivateKeyAlgorithm[0] = oidPublicKeyECDSA
	pkcs8Key.PrivateKeyAlgorithm[1] = oidNamedCurve
	pkcs8Key.PrivateKey = asn1Bytes

	pkcs8Bytes, err := asn1.Marshal(pkcs8Key)
	if err != nil {
		return nil, fmt.Errorf("error marshaling EC key to asn1 [%s]", err)
	}
	return pkcs8Bytes, nil
	/*
		return pem.EncodeToMemory(
			&pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: pkcs8Bytes,
			},
		), nil
	*/
}

func oidFromNamedCurve(curve elliptic.Curve) (asn1.ObjectIdentifier, bool) {
	switch curve {
	case elliptic.P224():
		return oidNamedCurveP224, true
	case elliptic.P256():
		return oidNamedCurveP256, true
	case elliptic.P384():
		return oidNamedCurveP384, true
	case elliptic.P521():
		return oidNamedCurveP521, true
	}
	return nil, false
}

func GenSerial(za []byte) string {
	hash := sha256.New()
	hash.Write(za)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func ValidateKey(dA []byte, P1 ecdsa.PublicKey, Pa []byte, ID, OU, Role string) error {
	c := elliptic.P256()
	n := elliptic.P256().Params().N

	d := new(big.Int).SetBytes(dA)
	priv := new(ecdsa.PrivateKey)
	priv.D = d
	priv.PublicKey.Curve = c
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(d.Bytes())
	X := priv.PublicKey.X
	Y := priv.PublicKey.Y
	X.Mod(X, n)
	Y.Mod(Y, n)

	var buffer bytes.Buffer
	buffer.Write([]byte(ID))
	buffer.Write(Pa)
	buffer.Write([]byte(OU))
	buffer.Write([]byte(Role))
	e := sha256.Sum256(buffer.Bytes())
	e0 := new(big.Int).SetBytes(e[0:15])
	e1 := new(big.Int).SetBytes(e[16:31])

	//get x1, y1, xa, ya
	puba, err := x509.ParsePKIXPublicKey(Pa)
	if err != nil {
		return err
	}
	pa := puba.(*ecdsa.PublicKey)

	xa, ya := c.ScalarMult(pa.X, pa.Y, e0.Bytes())
	x1, y1 := c.ScalarMult(P1.X, P1.Y, e1.Bytes())

	// x1 + xa ?= x
	x, y := c.Add(x1, y1, xa, ya)
	x.Mod(x, n)
	y.Mod(y, n)

	if x.Cmp(X) != 0 {
		return errors.New(("faield to pass key verification"))
	}
	return nil

}

func PrivateKeyToDER(d []byte, c elliptic.Curve) ([]byte, error) {
	k := new(ecdsa.PrivateKey)
	k.Curve = c
	k.D = new(big.Int).SetBytes(d)
	k.PublicKey.X, k.PublicKey.Y = c.ScalarBaseMult(d)

	// get the oid for the curve
	oidNamedCurve, ok := oidFromNamedCurve(k.Curve)
	if !ok {
		return nil, errors.New("unknown elliptic curve")
	}

	// based on https://golang.org/src/crypto/x509/sec1.go
	privateKeyBytes := k.D.Bytes()
	paddedPrivateKey := make([]byte, (k.Curve.Params().N.BitLen()+7)/8)
	copy(paddedPrivateKey[len(paddedPrivateKey)-len(privateKeyBytes):], privateKeyBytes)
	// omit NamedCurveOID for compatibility as it's optional
	asn1Bytes, err := asn1.Marshal(ecPrivateKey{
		Version:    1,
		PrivateKey: paddedPrivateKey,
		PublicKey:  asn1.BitString{Bytes: elliptic.Marshal(k.Curve, k.X, k.Y)},
	})

	if err != nil {
		return nil, fmt.Errorf("error marshaling EC key to asn1 [%s]", err)
	}

	var pkcs8Key pkcs8Info
	pkcs8Key.Version = 0
	pkcs8Key.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 2)
	pkcs8Key.PrivateKeyAlgorithm[0] = oidPublicKeyECDSA
	pkcs8Key.PrivateKeyAlgorithm[1] = oidNamedCurve
	pkcs8Key.PrivateKey = asn1Bytes

	pkcs8Bytes, err := asn1.Marshal(pkcs8Key)
	if err != nil {
		return nil, fmt.Errorf("error marshaling EC key to asn1 [%s]", err)
	}
	return pkcs8Bytes, nil
	/*
		return pem.EncodeToMemory(
			&pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: pkcs8Bytes,
			},
		), nil
	*/
}
