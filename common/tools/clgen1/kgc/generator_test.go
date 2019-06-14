/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package kgc_test

import (
	"crypto/ecdsa"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/fabric/common/tools/clgen/csp"
	"github.com/hyperledger/fabric/common/tools/clgen/kgc"
	"github.com/stretchr/testify/assert"
)

const (
	testKGCName            = "root"
	testOrganizationalUnit = "Hyperledger Fabric"
	testID                 = "Alice"
)

var testDir = filepath.Join(os.TempDir(), "kgc-test")

func TestNewKGC(t *testing.T) {
	kgcDir := filepath.Join(testDir, "kgc")
	rootKGC, err := kgc.NewKGC(kgcDir, testOrganizationalUnit, testKGCName)
	assert.NoError(t, err, "Error generating KGC")
	assert.NotNil(t, rootKGC, "Failed to return KGC")
	assert.NotNil(t, rootKGC.RawPub, "RawPub should not be empty")
	assert.IsType(t, &ecdsa.PrivateKey{}, rootKGC.MasterKey,
		"rootKGC.MasterKey should be type ecdsa.PrivateKey")
	assert.IsType(t, &ecdsa.PublicKey{}, rootKGC.MasterPub,
		"rootKGC.MasterPub should be type ecdsa.PublicKey")

	// check to make sure the root kgc public key was stored
	pemFile := filepath.Join(kgcDir, testKGCName+"-pubkey.pem")
	assert.Equal(t, true, checkForFile(pemFile),
		"Expected to find file "+pemFile)

	assert.NotEmpty(t, rootKGC.Name, "name cannot be empty.")
	assert.Equal(t, testKGCName, rootKGC.Name, "Failed to match Name")

	assert.NotEmpty(t, rootKGC.Organization, "Organization cannot be empty.")
	assert.Equal(t, testOrganizationalUnit, rootKGC.Organization, "Failed to match OrganizationalUnit")

	cleanup(testDir)
}

func TestKGCGenPartialKey(t *testing.T) {
	//	baseDir, ID string, XA *ecdsa.PublicKey, s *ecdsa.PrivateKey) (*PartialKey, error)

	kgcDir := filepath.Join(testDir, "kgc")
	certDir := filepath.Join(testDir, "certs")

	//create KGC
	rootKGC, err := kgc.NewKGC(kgcDir, testOrganizationalUnit, testKGCName)
	assert.NoError(t, err, "Error generating KGC")

	// generate x
	priv, _, err := csp.GeneratePrivateKey(certDir)
	assert.NoError(t, err, "Failed to generate client random x")

	// get XA
	ecPubKey, err := csp.GetECPublicKey(priv)
	assert.NoError(t, err, "Failed to generate XA")
	assert.NotNil(t, ecPubKey, "Failed to generate XA")

	//run
	partialK, err := rootKGC.KGCGenPartialKey(certDir, testID, ecPubKey)
	assert.NoError(t, err, "Failed to generate partial key")

	//check to make sure the partial key was stored
	pemFile := filepath.Join(certDir, testID+"-PA.pem")
	assert.Equal(t, true, checkForFile(pemFile),
		"Expected to find file "+pemFile)

	//load the pem file and compare
	rawPubKey, err := ioutil.ReadFile(pemFile)
	block, _ := pem.Decode(rawPubKey)
	assert.Equal(t, block.Bytes, partialK.PABytes(), "PA pem file not match")

	cleanup(testDir)
}

//func LoadKGCPublicKey(certPath string) (*ecdsa.PublicKey, []byte, error) {
func TestLoadKGCPublicKey(t *testing.T) {
	kgcDir := filepath.Join(testDir, "kgc")

	//create KGC
	rootKGC, err := kgc.NewKGC(kgcDir, testOrganizationalUnit, testKGCName)
	assert.NoError(t, err, "Error generating KGC")

	//run
	PubKey, raw, err := kgc.LoadKGCPublicKey(kgcDir)
	assert.NoError(t, err, "Error loading KGC public key")
	assert.Equal(t, PubKey, rootKGC.MasterPub, "KGC public key not match")
	assert.Equal(t, raw, rootKGC.RawPub, "KGC raw public key not match")

	cleanup(testDir)
}

func cleanup(dir string) {
	os.RemoveAll(dir)
}

func checkForFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
