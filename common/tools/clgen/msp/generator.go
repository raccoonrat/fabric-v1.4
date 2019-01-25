/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package msp

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/tools/clgen/ca"
	"github.com/hyperledger/fabric/common/tools/clgen/csp"
	"github.com/hyperledger/fabric/common/tools/clgen/kgc"
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
	// get keystore path
	keystore := filepath.Join(mspDir, "keystore")

	// generate x
	priv, _, err := csp.GeneratePrivateKey(keystore)
	if err != nil {
		return err
	}

	// get X
	ecPubKey, err := csp.GetECPublicKey(priv)
	if err != nil {
		return err
	}

	//kgc generate partial keys
	partialkey, err := signKGC.KGCGenPartialKey(filepath.Join(mspDir, "signcerts"),
		name, ecPubKey, signKGC.MasterKey)

	if err != nil {
		return err
	}

	//client generate final key
	err := csp.GenFinalKeyPair(keystore, priv, partialkey)
	if err != nil {
		return err
	}

	// write artifacts to MSP folders

	// the signing KGC pubkey goes into kgcpubs
	err = MasterPubExport(filepath.Join(baseDir, "kgcpubs", MasterPubFilename(signKGC.Name)), signKGC.RawPub)
	if err != nil {
		return err
	}
	// the TLS CA certificate goes into tlscacerts
	err = x509Export(filepath.Join(mspDir, "tlscacerts", x509Filename(tlsCA.Name)), tlsCA.SignCert)
	if err != nil {
		return err
	}

	// generate config.yaml if required
	if nodeOUs && nodeType == PEER {
		exportConfig(mspDir, "kgcpubs/"+x509Filename(signKGC.Name), true)
	}

	// the signing identity goes into admincerts.
	// This means that the signing identity
	// of this MSP is also an admin of this MSP
	// NOTE: the admincerts folder is going to be
	// cleared up anyway by copyAdminCert, but
	// we leave a valid admin for now for the sake
	// of unit tests
	err = MasterPubExport(filepath.Join(mspDir, "admincerts", MasterPubFilename(name)),
		partialkey.PartialPublickKey.Bytes())
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
		err = MasterPubExport(filepath.Join(baseDir, "kgcpubs", MasterPubFilename(signKGC.Name)), signKGC.RawPub)
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
		exportConfig(baseDir, "kgcpubs/"+x509Filename(signKGC.Name), true)
	}

	// create a throwaway PA to act as an admin PA
	// NOTE: the admincerts folder is going to be
	// cleared up anyway by copyAdminCert, but
	// we leave a valid admin for now for the sake
	// of unit tests
	factory.InitFactories(nil)
	bcsp := factory.GetDefault()
	priv, err := bcsp.KeyGen(&bccsp.ECDSAP256KeyGenOpts{Temporary: true})
	ecPubKey, err := csp.GetECPublicKey(priv)
	if err != nil {
		return err
	}
	_, err = signKGC.KGCGenPartialKey(filepath.Join(baseDir, "admincerts"), signKGC.Name,
		ecPubKey, signKGC.MasterKey)
	if err != nil {
		return err
	}

	return nil
}

func createFolderStructure(rootDir string, local bool) error {

	var folders []string
	// create admincerts, kgcpubkeys, keystore and signcerts folders
	folders = []string{
		filepath.Join(rootDir, "admincerts"),
		filepath.Join(rootDir, "kgcpubs"),
		filepath.Join(rootDir, "tlscacerts"),
	}
	if local {
		folders = append(folders, filepath.Join(rootDir, "keystore"),
			filepath.Join(rootDir, "signcerts"))
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
	return name + "-pubkey.pem"
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
