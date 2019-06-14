/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package msp

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/common/tools/clgen/ca"
	"github.com/hyperledger/fabric/common/tools/clgen/csp"
	"github.com/hyperledger/fabric/common/tools/clgen/kgc"
	idconfig "github.com/hyperledger/fabric/ibpcla/identity"
	fabricmsp "github.com/hyperledger/fabric/msp"
)

const (
	CLIENT = iota
	ORDERER
	PEER
)

const (
	CLIENTOU = "client"
	PEEROU   = "peer"
)

var nodeOUMap = map[int]string{
	CLIENT: CLIENTOU,
	PEER:   PEEROU,
}

func GenerateLocalMSP(baseDir, name string, sans []string, signKGC *kgc.KGC,
	tlsCA *ca.CA, nodeType int, nodeOUs bool) error {

	// create folder structure
	mspDir := filepath.Join(baseDir, "msp")
	tlsDir := filepath.Join(baseDir, "tls")

	err := createFolderStructure(mspDir, true)
	if err != nil {
		return err
	}

	err = os.MkdirAll(tlsDir, 0755)
	if err != nil {
		return err
	}

	/*
		Create the MSP identity artifacts
	*/

	// generate x
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	//kgc generate partial keys
	pa, za, err := signKGC.KGCGenPartialKey(name, &priv.PublicKey)
	if err != nil {
		return err
	}

	//client generate final key
	sk, err := csp.GenFinalKeyPair(name, priv, pa, za)
	if err != nil {
		return err
	}

	err = csp.ValidateKey(sk, signKGC.MasterKey.PublicKey, pa, name)
	if err != nil {
		return err
	}

	serial := csp.GenSerial(za)
	if err != nil {
		return err
	}

	var role string
	switch nodeType {
	case PEER:
		role = "peer"
	case ORDERER:
		role = "client"
	case CLIENT:
		if strings.Split(name, "@")[0] == "Admin" {
			role = "admin"
		} else {
			role = "member"
		}
	}
	IDConfig := &idconfig.IdConfig{
		Pk:                           pa,
		Sk:                           sk,
		Serial:                       serial,
		EnrollmentID:                 name,
		OrganizationalUnitIdentifier: signKGC.Organization,
		Role:                         role,
	}
	err = IDConfig.Store(filepath.Join(mspDir, "CLID", "IDconfig"))
	if err != nil {
		return err
	}

	// write artifacts to MSP folders

	// the signing KGC pubkey goes into kgcpubs
	err = ioutil.WriteFile(filepath.Join(mspDir, "kgcpubs", MasterPubFilename(signKGC.Name)), signKGC.RawPub, 0644)
	if err != nil {
		return err
	}
	// the TLS CA certificate goes into tlscacerts
	err = x509Export(filepath.Join(mspDir, "tlscacerts", x509Filename(tlsCA.Name)), tlsCA.SignCert)
	if err != nil {
		return err
	}

	// generate config.yaml if required
	EnableNode := nodeOUs && nodeType == PEER
	exportConfigID(mspDir, "kgcpubs/"+MasterPubFilename(signKGC.Name), EnableNode)

	// the signing identity goes into adminconfigs.
	// This means that the signing identity
	// of this MSP is also an admin of this MSP
	// NOTE: the adminconfigs folder is going to be
	// cleared up anyway by copyAdminCert, but
	// we leave a valid admin for now for the sake
	// of unit tests
	adminConfig := &idconfig.IdConfig{
		Pk:                           pa,
		EnrollmentID:                 name,
		OrganizationalUnitIdentifier: signKGC.Organization,
		Role:                         role,
	}
	err = adminConfig.Store(filepath.Join(mspDir, "CLID", "adminconfig"))
	if err != nil {
		return err
	}

	/*
		Generate the TLS artifacts in the TLS folder
	*/

	// generate private key
	tlsPrivKey, _, err := csp.GeneratePrivateKey(tlsDir)
	if err != nil {
		return err
	}
	// get public key
	tlsPubKey, err := csp.GetECPublicKey(tlsPrivKey)
	if err != nil {
		return err
	}
	// generate X509 certificate using TLS CA
	_, err = tlsCA.SignCertificate(filepath.Join(tlsDir),
		name, nil, sans, tlsPubKey, x509.KeyUsageDigitalSignature|x509.KeyUsageKeyEncipherment,
		[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth})
	if err != nil {
		return err
	}
	err = x509Export(filepath.Join(tlsDir, "ca.crt"), tlsCA.SignCert)
	if err != nil {
		return err
	}

	// rename the generated TLS X509 cert
	tlsFilePrefix := "server"
	if nodeType == CLIENT {
		tlsFilePrefix = "client"
	}
	err = os.Rename(filepath.Join(tlsDir, x509Filename(name)),
		filepath.Join(tlsDir, tlsFilePrefix+".crt"))
	if err != nil {
		return err
	}

	err = keyExport(tlsDir, filepath.Join(tlsDir, tlsFilePrefix+".key"), tlsPrivKey)
	if err != nil {
		return err
	}

	return nil
}

func GenerateVerifyingMSP(baseDir string, signKGC *kgc.KGC, tlsCA *ca.CA, nodeOUs bool) error {

	// create folder structure and write artifacts to proper locations
	err := createFolderStructure(baseDir, false)
	if err == nil {
		// the KGC Pubkeys goes into kgcpubs
		err = ioutil.WriteFile(filepath.Join(baseDir, "kgcpubs", MasterPubFilename(signKGC.Name)), signKGC.RawPub, 0644)
		if err != nil {
			return err
		}
		// the TLS CA certificate goes into tlscacerts
		err = x509Export(filepath.Join(baseDir, "tlscacerts", x509Filename(tlsCA.Name)), tlsCA.SignCert)
		if err != nil {
			return err
		}
	}

	// generate config.yaml if required
	if nodeOUs {
		exportConfig(baseDir, "kgcpubs/"+MasterPubFilename(signKGC.Name), true)
	}

	// create a throwaway PA to act as an admin PA
	// NOTE: the admincerts folder is going to be
	// cleared up anyway by copyAdminCert, but
	// we leave a valid admin for now for the sake
	// of unit tests
	/*
		factory.InitFactories(nil)
		bcsp := factory.GetDefault()
		priv, err := bcsp.KeyGen(&bccsp.ECDSAP256KeyGenOpts{Temporary: true})
		ecPubKey, err := csp.GetECPublicKey(priv)
	*/
	pri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}
	tempID := "TempAdmin"
	tempPA, _, err := signKGC.KGCGenPartialKey(tempID, &pri.PublicKey)
	if err != nil {
		return err
	}
	adminConfig := &idconfig.IdConfig{
		Pk:           tempPA,
		EnrollmentID: tempID,
	}
	err = adminConfig.Store(filepath.Join(baseDir, "CLID", "adminconfig"))
	if err != nil {
		return err
	}
	return nil
}

func createFolderStructure(rootDir string, local bool) error {

	var folders []string
	// create admincerts, kgcpubkeys, keystore and signcerts folders
	folders = []string{
		filepath.Join(rootDir, "CLID"),
		filepath.Join(rootDir, "kgcpubs"),
		filepath.Join(rootDir, "tlscacerts"),
	}
	if local {
	}

	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func x509Filename(name string) string {
	return name + "-cert.pem"
}

func x509Export(path string, cert *x509.Certificate) error {
	return pemExport(path, "CERTIFICATE", cert.Raw)
}

func MasterPubFilename(name string) string {
	return name + "-pubkey"
}

func MasterPubExport(path string, raw []byte) error {
	return pemExport(path, "ECC PUBLIC KEY", raw)
}

func keyExport(keystore, output string, key bccsp.Key) error {
	id := hex.EncodeToString(key.SKI())

	return os.Rename(filepath.Join(keystore, id+"_sk"), output)
}

func pemExport(path, pemType string, bytes []byte) error {
	//write pem out to file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, &pem.Block{Type: pemType, Bytes: bytes})
}

func exportConfig(mspDir, kgcFile string, enable bool) error {
	var config = &fabricmsp.Configuration{
		NodeOUs: &fabricmsp.NodeOUs{
			Enable: enable,
			ClientOUIdentifier: &fabricmsp.OrganizationalUnitIdentifiersConfiguration{
				Certificate:                  kgcFile,
				OrganizationalUnitIdentifier: CLIENTOU,
			},
			PeerOUIdentifier: &fabricmsp.OrganizationalUnitIdentifiersConfiguration{
				Certificate:                  kgcFile,
				OrganizationalUnitIdentifier: PEEROU,
			},
		},
	}

	configBytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(mspDir, "config.yaml"))
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(string(configBytes))

	return err
}

func exportConfigID(mspDir, kgcFile string, enable bool) error {
	var config = &fabricmsp.Configuration{}
	if enable {
		config.NodeOUs = &fabricmsp.NodeOUs{
			Enable: enable,
			ClientOUIdentifier: &fabricmsp.OrganizationalUnitIdentifiersConfiguration{
				Certificate:                  kgcFile,
				OrganizationalUnitIdentifier: CLIENTOU,
			},
			PeerOUIdentifier: &fabricmsp.OrganizationalUnitIdentifiersConfiguration{
				Certificate:                  kgcFile,
				OrganizationalUnitIdentifier: PEEROU,
			},
		}
	}

	configBytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(mspDir, "config.yaml"))
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(string(configBytes))

	return err
}
