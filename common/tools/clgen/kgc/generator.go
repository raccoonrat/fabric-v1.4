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
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/hyperledger/fabric/common/tools/clgen/csp"
)

type KGC struct {
	Name         string
	MasterKey    *ecdsa.PrivateKey
	MasterPub    *ecdsa.PublicKey
	RawPub       []byte
	Organization string
}

type PartialKey struct {
	PartialPublickKey *big.Int
	PartialPrivateKey *big.Int
}

// NewKGC creates an instance of KGC and saves the signing key pair in
// baseDir/name
func NewKGC(baseDir, org, name string) (*KGC, error) {

	var response error
	var kgc *KGC

	err := os.MkdirAll(baseDir, 0755)
	if err == nil {
		priv, err := csp.KGCGeneratePrivateKey(baseDir)
		response = err
		if err == nil {
			PubKey, raw, err := csp.KGCGetECPublicKey(priv, name, baseDir)
			response = err
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
	return kgc, response
}

// KGCGenPartialKey creates partial pk and sk based on a built-in template
// and saves it in baseDir/name
func (kgc *KGC) KGCGenPartialKey(baseDir, ID string, XA *ecdsa.PublicKey, s *ecdsa.PrivateKey) (*PartialKey, error) {

	var partialkey *PartialKey
	partialkey.PartialPublickKey, partialkey.PartialPrivateKey, err = KGCGenPartialKeyInternal(ID, XA, s)

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
	err = pem.Encode(PAFile, &pem.Block{Type: "ECC PUBLIC KEY", Bytes: partialkey.PartialPublickKey.Bytes()})
	PAFile.Close()
	if err != nil {
		return nil, err
	}

	return partialkey, nil
}

func KGCGenPartialKeyInternal(ID string, XA *ecdsa.PublicKey, s *ecdsa.PrivateKey) (*big.Int, *big.Int, error) {

	var buffer bytes.Buffer

	//get ecc base param n
	n := s.Curve.Params().N

	//y = rand()
	y, err := ecdsa.GenerateKey(s.Curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	//PA = XA + y*G
	//here we use x coordinate only
	PA := new(big.Int).Add(y.PublicKey.X, XA.X)
	PA.mod(PA, n)

	//e = hash(ID||PA)
	buffer.Write([]byte(ID))
	buffer.Write(PA.Bytes())
	e := sha256.Sum256(buffer)

	//e0=e[0:15], e1=e[16:32]

	e0 := new(big.Int).SetBytes(e[0:15])
	e1 := new(big.Int).SetBytes(e[16:32])

	//za = e0y + e1s
	e0.Mul(y.D, e0)
	e0.mod(e0, n)
	e1.Mul(s.D, e1)
	e1.mod(e1, n)
	za := new(big.Int).add(e0, e1)
	za.mod(za, n)

	return PA, za, nil
}

// default template for X509 subject
func subjectTemplate() pkix.Name {
	return pkix.Name{
		Country:  []string{"US"},
		Locality: []string{"San Francisco"},
		Province: []string{"California"},
	}
}

// Additional for X509 subject
func subjectTemplateAdditional(country, province, locality, orgUnit, streetAddress, postalCode string) pkix.Name {
	name := subjectTemplate()
	if len(country) >= 1 {
		name.Country = []string{country}
	}
	if len(province) >= 1 {
		name.Province = []string{province}
	}

	if len(locality) >= 1 {
		name.Locality = []string{locality}
	}
	if len(orgUnit) >= 1 {
		name.OrganizationalUnit = []string{orgUnit}
	}
	if len(streetAddress) >= 1 {
		name.StreetAddress = []string{streetAddress}
	}
	if len(postalCode) >= 1 {
		name.PostalCode = []string{postalCode}
	}
	return name
}

// generate a signed X509 certificate using ECDSA
func genCertificateECDSA(baseDir, name string, template, parent *x509.Certificate, pub *ecdsa.PublicKey,
	priv interface{}) (*x509.Certificate, error) {

	//create the x509 public cert
	certBytes, err := x509.CreateCertificate(rand.Reader, template, parent, pub, priv)
	if err != nil {
		return nil, err
	}

	//write cert out to file
	fileName := filepath.Join(baseDir, name+"-cert.pem")
	certFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	//pem encode the cert
	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	certFile.Close()
	if err != nil {
		return nil, err
	}

	x509Cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}
	return x509Cert, nil
}

// LoadCertificateECDSA load a ecdsa cert from a file in cert path
func LoadCertificateECDSA(certPath string) (*x509.Certificate, error) {
	var cert *x509.Certificate
	var err error

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".pem") {
			rawCert, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			block, _ := pem.Decode(rawCert)
			cert, err = utils.DERToX509Certificate(block.Bytes)
		}
		return nil
	}

	err = filepath.Walk(certPath, walkFunc)
	if err != nil {
		return nil, err
	}

	return cert, err
}
