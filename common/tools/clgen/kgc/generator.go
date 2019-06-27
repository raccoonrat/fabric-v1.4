/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package kgc

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric/common/tools/clgen/csp"
)

type KGC struct {
	Name         string
	MasterKey    *ecdsa.PrivateKey
	RawPub       []byte
	Organization string
}

type PPublicKey struct {
	X *big.Int
	Y *big.Int
}

type PartialKey struct {
	PartialPublicKey  *PPublicKey
	PartialPrivateKey *big.Int
}

func (PA *PartialKey) PABytes() []byte {

	var buffer bytes.Buffer
	buffer.Write(PA.PartialPublicKey.X.Bytes())
	buffer.Write(PA.PartialPublicKey.Y.Bytes())

	return buffer.Bytes()
}

// NewKGC creates an instance of KGC and saves the signing key pair in
// baseDir/name
func NewKGC(baseDir, org, name string) (*KGC, error) {

	var response error
	var kgc *KGC
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		return nil, err
	}

	priv, raw, err := csp.KGCGenerateMasterKey(baseDir)
	if err != nil {
		return nil, err
	}
	kgc = &KGC{
		Name:         name,
		MasterKey:    priv,
		RawPub:       raw,
		Organization: org,
	}
	return kgc, response
}

// KGCGenPartialKey creates partial pk and sk based on a built-in template
// and saves it in baseDir/name
func (kgc *KGC) KGCGenPartialKey(ID, role string, XA *ecdsa.PublicKey) ([]byte, []byte, error) {

	pa, za, err := KGCGenPartialKeyInternal(ID, kgc.Organization, role, XA, kgc.MasterKey)
	if err != nil {
		return nil, nil, err
	}

	return pa, za, nil
}

func KGCGenPartialKeyInternal(ID, OU, Role string, XA *ecdsa.PublicKey, s *ecdsa.PrivateKey) ([]byte, []byte, error) {

	var buffer bytes.Buffer

	//get ecc base param n
	c := s.Curve
	if c == nil {
		return nil, nil, errors.New("can not load curve params from master private key")
	}
	n := c.Params().N

	//y = rand()
	y, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	//PA = XA + y*G
	//here we use x coordinate only
	//no, have to use both
	PAx, PAy := c.Add(XA.X, XA.Y, y.PublicKey.X, y.PublicKey.Y)
	if PAx.Sign() == 0 && PAy.Sign() == 0 {
		return nil, nil, errors.New("invalid PA is generated")
	}

	PAx.Mod(PAx, n)
	PAy.Mod(PAy, n)

	var PA *ecdsa.PublicKey
	PA = XA
	PA.X.SetBytes(PAx.Bytes())
	PA.Y.SetBytes(PAy.Bytes())
	PABytes, err := x509.MarshalPKIXPublicKey(PA)
	//e = hash(ID||PA)
	buffer.Write([]byte(ID))
	buffer.Write(PABytes)
	buffer.Write([]byte(OU))
	buffer.Write([]byte(Role))
	e := sha256.Sum256(buffer.Bytes())

	//e0=e[0:15], e1=e[16:31]

	e0 := new(big.Int).SetBytes(e[0:15])
	e1 := new(big.Int).SetBytes(e[16:31])

	//za = e0y + e1s
	e0.Mul(y.D, e0)
	e1.Mul(s.D, e1)
	za := new(big.Int).Add(e0, e1)
	za.Mod(za, n)

	return PABytes, za.Bytes(), nil
}

// LoadKGCPublicKey load a ecdsa public key from a file in cert path
func LoadKGCPublicKey(certPath string) (*ecdsa.PublicKey, []byte, error) {
	var Pub *ecdsa.PublicKey
	var raw []byte
	var err error

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".pem") {
			rawPubKey, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			block, _ := pem.Decode(rawPubKey)
			ecPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return err
			}
			Pub = ecPubKey.(*ecdsa.PublicKey)
			raw = block.Bytes
		}
		return nil
	}

	err = filepath.Walk(certPath, walkFunc)
	if err != nil {
		return nil, nil, err
	}

	return Pub, raw, err
}
