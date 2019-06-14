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
	MasterPub    *ecdsa.PublicKey
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
	if err == nil {
		bpriv, err := csp.KGCGeneratePrivateKey(baseDir)
		response = err
		if err == nil {
			PubKey, raw, err := csp.KGCGetECPublicKey(bpriv, name, baseDir)
			response = err
			if err == nil {
				priv, err := csp.LoadCLPrivateKey(baseDir, bpriv.SKI())
				if err == nil {
					kgc = &KGC{
						Name:         name,
						MasterKey:    priv,
						MasterPub:    PubKey,
						RawPub:       raw,
						Organization: org,
					}
				}
			}
		}
	}
	return kgc, response
}

// KGCGenPartialKey creates partial pk and sk based on a built-in template
// and saves it in baseDir/name
func (kgc *KGC) KGCGenPartialKey(baseDir, ID string, XA *ecdsa.PublicKey) (*PartialKey, error) {

	var partialkey *PartialKey
	pa, za, err := KGCGenPartialKeyInternal(ID, XA, kgc.MasterKey)
	if err == nil {
		partialkey = &PartialKey{
			PartialPublicKey:  pa,
			PartialPrivateKey: za,
		}
	}

	if err != nil {
		return nil, err
	}

	//write Partial Public Key to file
	fileName := filepath.Join(baseDir, ID+"-PA.pem")
	PAFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	//pem encode the public key
	//PA = PAx||PAy
	err = pem.Encode(PAFile, &pem.Block{Type: "PUBLIC KEY", Bytes: partialkey.PABytes()})
	PAFile.Close()
	if err != nil {
		return nil, err
	}

	//to load PA
	/*
		rawPubKey, err := ioutil.ReadFile(fileName)
		block, _ := pem.Decode(rawPubKey)
		fmt.Println("PA:" + hex.EncodeToString(block.Bytes))
		PAx := new(big.Int).SetBytes(block.Bytes[0:15])
		PAy := new(big.Int).SetBytes(block.Bytes[16:31])
	*/

	return partialkey, nil
}

func KGCGenPartialKeyInternal(ID string, XA *ecdsa.PublicKey, s *ecdsa.PrivateKey) (*PPublicKey, *big.Int, error) {

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

	//e = hash(ID||PA)
	buffer.Write([]byte(ID))
	buffer.Write(PAx.Bytes())
	buffer.Write(PAy.Bytes())
	e := sha256.Sum256(buffer.Bytes())

	//e0=e[0:15], e1=e[16:31]

	e0 := new(big.Int).SetBytes(e[0:15])
	e1 := new(big.Int).SetBytes(e[16:31])

	//za = e0y + e1s
	e0.Mul(y.D, e0)
	e1.Mul(s.D, e1)
	za := new(big.Int).Add(e0, e1)
	za.Mod(za, n)

	return &PPublicKey{X: PAx, Y: PAy}, za, nil
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
