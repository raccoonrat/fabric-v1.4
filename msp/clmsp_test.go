/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0

*/

package msp

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/core/config/configtest"
	"github.com/hyperledger/fabric/protos/msp"
	m "github.com/hyperledger/fabric/protos/msp"
	"github.com/stretchr/testify/assert"
)

const SampleOrg = "peer0.org1.example.com"

func CLgetconfig() (*msp.CLMSPConfig, error) {
	mspDir, err := configtest.GetDevCLMspDir()
	conf, err = GetLocalCLMspConfig(mspDir, nil, SampleOrg)
	clconfig := &msp.CLMSPConfig{}
	err = proto.Unmarshal(conf.Config, clconfig)
	if err != nil {
		return nil, err
	}
	return clconfig, nil
}

func TestCLMSPNormal(t *testing.T) {
	mspDir, err := configtest.GetDevCLMspDir()
	if err != nil {
		fmt.Printf("Errog getting DevMspDir: %s", err)
		os.Exit(-1)
	}

	conf1, err := GetLocalCLMspConfig(mspDir, nil, SampleOrg)
	fmt.Println("-----------------")
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMspCL, err = newIBPCLAMsp(MSPv1_3)
	if err != nil {
		fmt.Printf("Constructor for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMspBadCL, err = newIBPCLAMsp(MSPv1_3)
	if err != nil {
		fmt.Printf("Constructor for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMspCLV11, err = newIBPCLAMsp(MSPv1_1)
	if err != nil {
		fmt.Printf("Constructor for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMspCL.Setup(conf1)
	if err != nil {
		fmt.Printf("Setup for V1.3 msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMspBadCL.Setup(conf1)
	if err != nil {
		fmt.Printf("Setup for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMspCLV11.Setup(conf1)
	if err != nil {
		fmt.Printf("Setup for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	mspMgrCL = NewMSPManager()
	err = mspMgrCL.Setup([]MSP{localMspCL})
	if err != nil {
		fmt.Printf("Setup for msp manager should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	id, err := localMspCL.GetIdentifier()
	if err != nil {
		fmt.Println("Failed obtaining identifier for localMSP")
		os.Exit(-1)
	}

	msps, err := mspMgrCL.GetMSPs()
	if err != nil {
		fmt.Println("Failed obtaining MSPs from MSP manager")
		os.Exit(-1)
	}

	if msps[id] == nil {
		fmt.Println("Couldn't find localMSP in MSP manager")
		os.Exit(-1)
	}

	_, err = localMspCL.(*clmsp).getSigningIdentityFromConf(nil)
	assert.Error(t, err)
}
func TestCLMSPParsers(t *testing.T) {

	mspDir, err := configtest.GetDevCLMspDir()
	if err != nil {
		fmt.Printf("Errog getting DevMspDir: %s", err)
		os.Exit(-1)
	}

	conf, err := GetLocalCLMspConfig(mspDir, nil, SampleOrg)
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	clconf := &m.CLMSPConfig{}
	err = proto.Unmarshal(conf.Config, clconf)
	assert.NoError(t, err)

	sigid := &msp.CLMSPSignerConfig{Sk: clconf.CLSigningIdentity.Sk, PA: clconf.CLSigningIdentity.PA}
	_, err = localMspCL.(*clmsp).getSigningIdentityFromConf(sigid)
	assert.NoError(t, err)

	sigid = &msp.CLMSPSignerConfig{Sk: nil, PA: clconf.CLSigningIdentity.PA}
	_, err = localMspCL.(*clmsp).getSigningIdentityFromConf(sigid)
	assert.Error(t, err)

	sigid = &msp.CLMSPSignerConfig{Sk: clconf.CLSigningIdentity.PA, PA: nil}
	_, err = localMspCL.(*clmsp).getSigningIdentityFromConf(sigid)
	assert.Error(t, err)

	_, err = localMspCL.(*clmsp).getSigningIdentityFromConf(nil)
	assert.Error(t, err)
}

func TestCLMSPSetupNoCryptoConf(t *testing.T) {
	mspDir, err := configtest.GetDevCLMspDir()
	if err != nil {
		fmt.Printf("Errog getting DevMspDir: %s", err)
		os.Exit(-1)
	}

	conf, err := GetLocalCLMspConfig(mspDir, nil, SampleOrg)
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	mspconf := &msp.CLMSPConfig{}
	err = proto.Unmarshal(conf.Config, mspconf)
	assert.NoError(t, err)

	// here we test the case of an MSP configuration
	// where the hash function to be used to obtain
	// the identity identifier is unspecified - a
	// sane default should be picked
	mspconf.CryptoConfig.IdentityIdentifierHashFunction = ""
	b, err := proto.Marshal(mspconf)
	assert.NoError(t, err)
	conf.Config = b
	newmsp, err := newIBPCLAMsp(MSPv1_3)
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)
	// here we test the case of an MSP configuration
	// where the hash function to be used to compute
	// signatures is unspecified - a sane default
	// should be picked
	mspconf.CryptoConfig.SignatureHashFamily = ""
	b, err = proto.Marshal(mspconf)
	assert.NoError(t, err)
	conf.Config = b
	newmsp, err = newIBPCLAMsp(MSPv1_3)
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)

	// here we test the case of an MSP configuration
	// that has NO crypto configuration specified;
	// the code will use appropriate defaults
	mspconf.CryptoConfig = nil
	b, err = proto.Marshal(mspconf)
	assert.NoError(t, err)
	conf.Config = b
	newmsp, err = newIBPCLAMsp(MSPv1_3)
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)
}

func TestCLGetters(t *testing.T) {
	typ := localMspCL.GetType()
	assert.Equal(t, typ, IBPCLA)
	assert.NotNil(t, localMsp.GetTLSRootCerts())
	assert.NotNil(t, localMsp.GetTLSIntermediateCerts())
}

func TestCLMSPSetupBad(t *testing.T) {
	_, err := GetLocalCLMspConfig("barf", nil, SampleOrg)
	if err == nil {
		t.Fatalf("Setup should have failed on an invalid config file")
		return
	}

	mgr := NewMSPManager()
	err = mgr.Setup(nil)
	assert.NoError(t, err)
	err = mgr.Setup([]MSP{})
	assert.NoError(t, err)

	// Setup with nil config
	err = localMspBadCL.Setup(nil)
	assert.Error(t, err)

	// Setup with incorrect MSP type
	conf := &msp.MSPConfig{Type: 1234, Config: nil}
	err = localMspBadCL.Setup(conf)
	assert.Error(t, err)

	// Setup with bad IBPCLA Config bytes
	conf = &msp.MSPConfig{Type: int32(IBPCLA), Config: []byte("barf")}
	err = localMspBadCL.Setup(conf)
	assert.Error(t, err)

	clconfig, err := CLgetconfig()
	assert.NoError(t, err)

	// Create MSP config with bad KGCPubs
	var KGCPubByte [][]byte
	KGCPubByte = append(KGCPubByte, []byte("barf"))
	clconfig.KGCPubs = KGCPubByte
	clConfigBytes, err := proto.Marshal(clconfig)
	assert.NoError(t, err)
	conf.Config = clConfigBytes
	err = localMspBadCL.Setup(conf)
	assert.Error(t, err)

	// Create MSP config with bad Name
	clconfig, err = CLgetconfig()
	assert.NoError(t, err)
	clconfig.Name = ""
	clConfigBytes, err = proto.Marshal(clconfig)
	assert.NoError(t, err)
	conf.Config = clConfigBytes
	err = localMspBadCL.Setup(conf)
	assert.Error(t, err)

	// Create MSP config with bad Admins
	/*
		clconfig, err = CLgetconfig()
		assert.NoError(t, err)
		var AdminByte [][]byte
		AdminByte = append(AdminByte, []byte("barf"))
		clconfig.Admins = AdminByte
		clConfigBytes, err = proto.Marshal(clconfig)
		assert.NoError(t, err)
		conf.Config = clConfigBytes
		err = localMspBadCL.Setup(conf)
		assert.Error(t, err)
	*/
}

func TestCLDoubleSetup(t *testing.T) {
	// note that we've already called setup once on this
	err := mspMgrCL.Setup(nil)
	assert.NoError(t, err)
}

func TestCLGetIdentities(t *testing.T) {
	_, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetDefaultSigningIdentity failed with err %s", err)
		return
	}
}

func TestCLDeserializeIdentityFails(t *testing.T) {
	_, err := localMspCL.DeserializeIdentity([]byte("barf"))
	assert.Error(t, err)

	id := &msp.SerializedIdentity{Mspid: SampleOrg, IdBytes: []byte("barfr")}
	b, err := proto.Marshal(id)
	assert.NoError(t, err)
	_, err = localMspCL.DeserializeIdentity(b)
	assert.Error(t, err)

	id = &msp.SerializedIdentity{Mspid: SampleOrg, IdBytes: []byte(notACert)}
	b, err = proto.Marshal(id)
	assert.NoError(t, err)
	_, err = localMspCL.DeserializeIdentity(b)
	assert.Error(t, err)
}

func TestCLGetSigningIdentityFromVerifyingMSP(t *testing.T) {
	mspDir, err := configtest.GetDevCLMspDir()
	if err != nil {
		fmt.Printf("Errog getting DevMspDir: %s", err)
		os.Exit(-1)
	}

	conf, err := GetVerifyingMspConfig(mspDir, SampleOrg, ProviderTypeToString(IBPCLA))
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead. ", err)
		os.Exit(-1)
	}

	newmsp, err := newIBPCLAMsp(MSPv1_3)
	assert.NoError(t, err)
	err = newmsp.Setup(conf)
	assert.NoError(t, err)

	_, err = newmsp.GetDefaultSigningIdentity()
	assert.Error(t, err)
	_, err = newmsp.GetSigningIdentity(nil)
	assert.Error(t, err)
}

func TestCLValidateDefaultSigningIdentity(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	err = localMspCL.Validate(id.GetPublicVersion())
	assert.NoError(t, err)
}

func TestCLSerializeIdentities(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	idBack, err := localMspCL.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded, got err %s", err)
		return
	}

	err = localMspCL.Validate(idBack)
	if err != nil {
		t.Fatalf("The identity should be valid, got err %s", err)
		return
	}

	if !reflect.DeepEqual(id.GetPublicVersion(), idBack) {
		t.Fatalf("Identities should be equal (%s) (%s)", id, idBack)
		return
	}
}

func TestCLIsWellFormed(t *testing.T) {
	mspMgrCL := NewMSPManager()

	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	sId := &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sId)
	assert.NoError(t, err)

	// An MSP Manager without any MSPs should not recognize the identity since
	// no providers are registered
	err = mspMgrCL.IsWellFormed(sId)
	assert.Error(t, err)
	assert.Equal(t, "no MSP provider recognizes the identity", err.Error())

	// Add the MSP to the MSP Manager
	mspMgrCL.Setup([]MSP{localMspCL})

	err = localMspCL.IsWellFormed(sId)
	assert.NoError(t, err)
	err = mspMgrCL.IsWellFormed(sId)
	assert.NoError(t, err)

	sId.IdBytes = append(sId.IdBytes, 1)
	err = localMspCL.IsWellFormed(sId)
	assert.Error(t, err)

	err = mspMgrCL.IsWellFormed(sId)
	assert.Error(t, err)
	assert.Equal(t, "no MSP provider recognizes the identity", err.Error())
}

func TestCLValidateKGCIdentity(t *testing.T) {
	//currently Validate always success

	//kgcID := getCLIdentity(t, CLKGCPubs)
	//err := localMspCL.Validate(kgcID)
	//assert.Error(t, err)
}

/*
func TestCLBadAdminIdentity(t *testing.T) {
	conf, err := GetLocalCLMspConfig("testdata/badadmincl", nil, SampleOrg)
	assert.NoError(t, err)

	thisMSP, err := newIBPCLAMsp(MSPv1_3)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join("testdata/badadmincl", "keystore"), true)
	assert.NoError(t, err)
	csp, err := clbccsp.New(ks)
	assert.NoError(t, err)
	thisMSP.(*clmsp).csp = csp

	err = thisMSP.Setup(conf)
	//Validate always success
	//assert.Error(t, err)
}
*/

func TestCLValidateAdminIdentity(t *testing.T) {
	//currently Validate always success
	/*
		caID := getIdentity(t, admincerts)

		err := localMsp.Validate(caID)
		assert.NoError(t, err)
	*/
}

func TestCLSerializeIdentitiesWithWrongMSP(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	sid := &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sid)
	assert.NoError(t, err)

	sid.Mspid += "BARF"

	serializedID, err = proto.Marshal(sid)
	assert.NoError(t, err)

	_, err = localMspCL.DeserializeIdentity(serializedID)
	assert.Error(t, err)
}

func TestCLSerializeIdentitiesWithMSPManager(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	_, err = mspMgrCL.DeserializeIdentity(serializedID)
	assert.NoError(t, err)

	sid := &msp.SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sid)
	assert.NoError(t, err)

	sid.Mspid += "BARF"

	serializedID, err = proto.Marshal(sid)
	assert.NoError(t, err)

	_, err = mspMgrCL.DeserializeIdentity(serializedID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("MSP %s is unknown", sid.Mspid))

	_, err = mspMgrCL.DeserializeIdentity([]byte("barf"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not deserialize")
}

func TestCLIdentitiesGetters(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	idid := id.GetIdentifier()
	assert.NotNil(t, idid)
	mspid := id.GetMSPIdentifier()
	assert.NotNil(t, mspid)
	assert.False(t, id.Anonymous())
}

func TestCLSignAndVerify(t *testing.T) {
	sid, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	serializedID, err := sid.Serialize()
	assert.NoError(t, err)

	sidBack, err := localMspCL.DeserializeIdentity(serializedID)
	assert.NoError(t, err)

	msg := []byte("foo")
	sig, err := sid.Sign(msg)
	assert.NoError(t, err)

	err = sid.Verify(msg, sig)
	assert.NoError(t, err)
	fmt.Println("-----------1")

	err = sidBack.Verify(msg, sig)
	assert.NoError(t, err)
	fmt.Println("-----------2")

	err = sid.Verify(msg[1:], sig)
	assert.Error(t, err)
	fmt.Println("-----------3")

	err = sid.Verify(msg, sig[1:])
	assert.Error(t, err)
}

func TestCLSignAndVerifyFailures(t *testing.T) {
	msg := []byte("foo")

	id, err := localMspBadCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	hash := id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily
	id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily = "barf"

	sig, err := id.Sign(msg)
	assert.Error(t, err)

	id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily = hash

	sig, err = id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily = "barf"

	err = id.Verify(msg, sig)
	assert.Error(t, err)

	id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily = hash
}

func TestCLSignAndVerifyOtherHash(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	hash := id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily
	id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily = bccsp.SHA3

	msg := []byte("foo")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	assert.NoError(t, err)

	id.(*clsigningidentity).msp.cryptoConfig.SignatureHashFamily = hash
}

func TestCLSignAndVerify_longMessage(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded")
		return
	}

	idBack, err := localMspCL.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded")
		return
	}

	msg := []byte("ABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFG")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = idBack.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}
}

func TestCLGetOU(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	assert.Equal(t, "COP", id.GetOrganizationalUnits()[0].OrganizationalUnitIdentifier)
}

func TestCLGetOUFail(t *testing.T) {
	id, err := localMspBadCL.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	ouTmp := make(map[string][][]byte)
	id.(*clsigningidentity).msp.ouIdentifiers = ouTmp
	ou := id.GetOrganizationalUnits()
	assert.Nil(t, ou)
}

/*
func TestCLKGCIdentifierComputation(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	chain, err := localMsp.(*bccspmsp).getCertificationChain(id.GetPublicVersion())
	assert.NoError(t, err)

	// Hash the chain
	// Use the hash of the identity's certificate as id in the IdentityIdentifier
	hashOpt, err := bccsp.GetHashOpt(localMsp.(*bccspmsp).cryptoConfig.IdentityIdentifierHashFunction)
	assert.NoError(t, err)

	hf, err := localMsp.(*bccspmsp).bccsp.GetHash(hashOpt)
	assert.NoError(t, err)
	// Skipping first cert because it belongs to the identity
	for i := 1; i < len(chain); i++ {
		hf.Write(chain[i].Raw)
	}
	sum := hf.Sum(nil)

	assert.Equal(t, sum, id.GetOrganizationalUnits()[0].CertifiersIdentifier)
}
*/
func TestCLOUPolicyPrincipal(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	cid, err := localMspCL.(*clmsp).getPAIdentifier(id.GetPublicVersion())
	assert.NoError(t, err)

	ou := &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "peer0.org1.example.com",
		CertifiersIdentifier:         cid,
	}
	bytes, err := proto.Marshal(ou)
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestCLOUPolicyPrincipalBadPrincipal(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               []byte("barf"),
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLOUPolicyPrincipalBadMSPID(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	cid, err := localMspCL.(*clmsp).getPAIdentifier(id.GetPublicVersion())
	assert.NoError(t, err)

	ou := &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "SampleOrgbarfbarf",
		CertifiersIdentifier:         cid,
	}
	bytes, err := proto.Marshal(ou)
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLOUPolicyPrincipalBadPath(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	ou := &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "peer0.org1.example.com",
		CertifiersIdentifier:         nil,
	}
	bytes, err := proto.Marshal(ou)
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)

	ou = &msp.OrganizationUnit{
		OrganizationalUnitIdentifier: "COP",
		MspIdentifier:                "peer0.org1.example.com",
		CertifiersIdentifier:         []byte{0, 1, 2, 3, 4},
	}
	bytes, err = proto.Marshal(ou)
	assert.NoError(t, err)

	principal = &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ORGANIZATION_UNIT,
		Principal:               bytes,
	}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLPolicyPrincipalBogusType(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: 35, MspIdentifier: "SampleOrg"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: 35,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLPolicyPrincipalBogusRole(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: 35, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLPolicyPrincipalWrongMSPID(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "SampleOrgBARFBARF"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	//currently do not check ID
	assert.NoError(t, err)
}

func TestCLMemberPolicyPrincipal(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestCLAdminPolicyPrincipal(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

/*
// Combine one or more MSPPrincipals into a MSPPrincipal of type
// MSPPrincipal_COMBINED.
func createCombinedPrincipal(principals ...*msp.MSPPrincipal) (*msp.MSPPrincipal, error) {
	if len(principals) == 0 {
		return nil, errors.New("no principals in CombinedPrincipal")
	}
	var principalsArray []*msp.MSPPrincipal
	for _, principal := range principals {
		principalsArray = append(principalsArray, principal)
	}
	combinedPrincipal := &msp.CombinedPrincipal{Principals: principalsArray}
	combinedPrincipalBytes, err := proto.Marshal(combinedPrincipal)
	if err != nil {
		return nil, err
	}
	principalsCombined := &msp.MSPPrincipal{PrincipalClassification: msp.MSPPrincipal_COMBINED, Principal: combinedPrincipalBytes}
	return principalsCombined, nil
}
*/

func TestCLMultilevelAdminAndMemberPolicyPrincipal(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	adminPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	memberPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	adminPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               adminPrincipalBytes}

	memberPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               memberPrincipalBytes}

	// CombinedPrincipal with Admin and Member principals
	levelOneCombinedPrincipal, err := createCombinedPrincipal(adminPrincipal, memberPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)

	// Nested CombinedPrincipal
	levelTwoCombinedPrincipal, err := createCombinedPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)

	// Double nested CombinedPrincipal
	levelThreeCombinedPrincipal, err := createCombinedPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelThreeCombinedPrincipal)
	assert.NoError(t, err)
}

func TestCLMultilevelAdminAndMemberPolicyPrincipalPreV12(t *testing.T) {
	id, err := localMspCLV11.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	adminPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	memberPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	adminPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               adminPrincipalBytes}

	memberPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               memberPrincipalBytes}

	// CombinedPrincipal with Admin and Member principals
	levelOneCombinedPrincipal, err := createCombinedPrincipal(adminPrincipal, memberPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelOneCombinedPrincipal)
	assert.Error(t, err)

	// Nested CombinedPrincipal
	levelTwoCombinedPrincipal, err := createCombinedPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelTwoCombinedPrincipal)
	assert.Error(t, err)

	// Double nested CombinedPrincipal
	levelThreeCombinedPrincipal, err := createCombinedPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelThreeCombinedPrincipal)
	assert.Error(t, err)
}

func TestCLAdminPolicyPrincipalFails(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               principalBytes}

	// remove the admin so validation will fail
	localMspCL.(*clmsp).admins = make([]clidentity, 0)

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLMultilevelAdminAndMemberPolicyPrincipalFails(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	adminPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	memberPrincipalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: "peer0.org1.example.com"})
	assert.NoError(t, err)

	adminPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               adminPrincipalBytes}

	memberPrincipal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ROLE,
		Principal:               memberPrincipalBytes}

	// remove the admin so validation will fail
	localMspCL.(*clmsp).admins = make([]clidentity, 0)

	// CombinedPrincipal with Admin and Member principals
	levelOneCombinedPrincipal, err := createCombinedPrincipal(adminPrincipal, memberPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelOneCombinedPrincipal)
	assert.Error(t, err)

	// Nested CombinedPrincipal
	levelTwoCombinedPrincipal, err := createCombinedPrincipal(levelOneCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelTwoCombinedPrincipal)
	assert.Error(t, err)

	// Double nested CombinedPrincipal
	levelThreeCombinedPrincipal, err := createCombinedPrincipal(levelTwoCombinedPrincipal)
	assert.NoError(t, err)
	err = id.SatisfiesPrincipal(levelThreeCombinedPrincipal)
	assert.Error(t, err)
}

/*
func TestIdentityExpiresAt(t *testing.T) {
	thisMSP := getLocalMSP(t, "testdata/expiration")
	assert.NotNil(t, thisMSP)
	si, err := thisMSP.GetDefaultSigningIdentity()
	assert.NoError(t, err)
	expirationDate := si.GetPublicVersion().ExpiresAt()
	assert.Equal(t, time.Date(2027, 8, 17, 12, 19, 48, 0, time.UTC), expirationDate)
}

func TestIdentityExpired(t *testing.T) {
	expiredCertsDir := "testdata/expired"
	conf, err := GetLocalMspConfig(expiredCertsDir, nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(MSPv1_0)
	assert.NoError(t, err)

	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(expiredCertsDir, "keystore"), true)
	assert.NoError(t, err)

	csp, err := sw.NewWithParams(256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	err = thisMSP.Setup(conf)
	if err != nil {
		assert.Contains(t, err.Error(), "signing identity expired")
	} else {
		t.Fatal("Should have failed when loading expired certs")
	}
}
*/

func TestCLIdentityPolicyPrincipal(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	idSerialized, err := id.Serialize()
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_IDENTITY,
		Principal:               idSerialized}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestCLIdentityPolicyPrincipalBadBytes(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_IDENTITY,
		Principal:               []byte("barf")}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

/*
func TestCLMSPOus(t *testing.T) {
	// Set the OUIdentifiers
	backup := localMspCL.(*clmsp).ouIdentifiers
	defer func() { localMspCL.(*clmsp).ouIdentifiers = backup }()
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	localMspCL.(*clmsp).ouIdentifiers = map[string][][]byte{
		"COP": {id.GetOrganizationalUnits()[0].CertifiersIdentifier},
	}
	assert.NoError(t, localMspCL.Validate(id.GetPublicVersion()))

	localMspCL.(*clmsp).ouIdentifiers = map[string][][]byte{
		"COP2": {id.GetOrganizationalUnits()[0].CertifiersIdentifier},
	}
	assert.Error(t, localMspCL.Validate(id.GetPublicVersion()))

	localMspCL.(*clmsp).ouIdentifiers = map[string][][]byte{
		"COP": {{0, 1, 2, 3, 4}},
	}
	assert.Error(t, localMspCL.Validate(id.GetPublicVersion()))
}
*/

const otherPA = `-----BEGIN PUBLIC KEY-----
74slS6oQXNHIjMai3UHNnW628/u56++ZRa5IfAU4sWJpjabG+RCWTObfrQJHf6/B
69Wu5ggiw+Em6EBfcMNHEw==
-----END PUBLIC KEY-----
`

func TestCLIdentityPolicyPrincipalFails(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	sid, err := NewSerializedclIdentity("peer0.org1.example.com", SampleOrg, []byte(otherPA))
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_IDENTITY,
		Principal:               sid}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

var localMspCL MSP
var localMspBadCL MSP
var localMspCLV11 MSP
var mspMgrCL MSPManager

/*

func getIdentity(t *testing.T, path string) Identity {
	mspDir, err := configtest.GetDevMspDir()
	assert.NoError(t, err)

	pems, err := getPemMaterialFromDir(filepath.Join(mspDir, path))
	assert.NoError(t, err)

	id, _, err := localMsp.(*bccspmsp).getIdentityFromConf(pems[0])
	assert.NoError(t, err)

	return id
}

func getLocalMSPWithVersionAndError(t *testing.T, dir string, version MSPVersion) (MSP, error) {
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(version)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)
	csp, err := sw.NewWithParams(256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	return thisMSP, thisMSP.Setup(conf)
}

func getLocalMSP(t *testing.T, dir string) MSP {
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(MSPv1_0)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)
	csp, err := sw.NewWithParams(256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	err = thisMSP.Setup(conf)
	assert.NoError(t, err)

	return thisMSP
}

func getLocalMSPWithVersion(t *testing.T, dir string, version MSPVersion) MSP {
	conf, err := GetLocalMspConfig(dir, nil, "SampleOrg")
	assert.NoError(t, err)

	thisMSP, err := newBccspMsp(version)
	assert.NoError(t, err)
	ks, err := sw.NewFileBasedKeyStore(nil, filepath.Join(dir, "keystore"), true)
	assert.NoError(t, err)
	csp, err := sw.NewWithParams(256, "SHA2", ks)
	assert.NoError(t, err)
	thisMSP.(*bccspmsp).bccsp = csp

	err = thisMSP.Setup(conf)
	assert.NoError(t, err)

	return thisMSP
}
*/

func TestCLCollectEmptyCombinedPrincipal(t *testing.T) {
	var principalsArray []*msp.MSPPrincipal
	combinedPrincipal := &msp.CombinedPrincipal{Principals: principalsArray}
	combinedPrincipalBytes, err := proto.Marshal(combinedPrincipal)
	assert.NoError(t, err, "Error marshalling empty combined principal")
	principalsCombined := &msp.MSPPrincipal{PrincipalClassification: msp.MSPPrincipal_COMBINED, Principal: combinedPrincipalBytes}
	_, err = collectPrincipals(principalsCombined, MSPv1_3)
	assert.Error(t, err)
}

func TestCLCollectPrincipalContainingEmptyCombinedPrincipal(t *testing.T) {
	var principalsArray []*msp.MSPPrincipal
	combinedPrincipal := &msp.CombinedPrincipal{Principals: principalsArray}
	combinedPrincipalBytes, err := proto.Marshal(combinedPrincipal)
	assert.NoError(t, err, "Error marshalling empty combined principal")
	emptyPrincipal := &msp.MSPPrincipal{PrincipalClassification: msp.MSPPrincipal_COMBINED, Principal: combinedPrincipalBytes}
	levelOneCombinedPrincipal, err := createCombinedPrincipal(emptyPrincipal)
	assert.NoError(t, err)
	_, err = collectPrincipals(levelOneCombinedPrincipal, MSPv1_3)
	assert.Error(t, err)
}

/*
func TestCLMSPIdentityIdentifier(t *testing.T) {
	// testdata/mspid
	// 1) a key and a signcert (used to populate the default signing identity) with the cert having a HighS signature
	thisMSP := getLocalMSP(t, "testdata/mspid")

	id, err := thisMSP.GetDefaultSigningIdentity()
	assert.NoError(t, err)
	err = id.Validate()
	assert.NoError(t, err)

	// Check that the identity identifier is computed with the respect to the lowS signature

	idid := id.GetIdentifier()
	assert.NotNil(t, idid)

	// Load and parse cacaert and signcert from folder
	pems, err := getPemMaterialFromDir("testdata/mspid/cacerts")
	assert.NoError(t, err)
	bl, _ := pem.Decode(pems[0])
	assert.NotNil(t, bl)
	caCertFromFile, err := x509.ParseCertificate(bl.Bytes)
	assert.NoError(t, err)

	pems, err = getPemMaterialFromDir("testdata/mspid/signcerts")
	assert.NoError(t, err)
	bl, _ = pem.Decode(pems[0])
	assert.NotNil(t, bl)
	certFromFile, err := x509.ParseCertificate(bl.Bytes)
	assert.NoError(t, err)
	// Check that the certificates' raws are different, meaning that the identity has been sanitised
	assert.NotEqual(t, certFromFile.Raw, id.(*signingidentity).cert)

	// Check that certFromFile is in HighS
	_, S, err := utils.UnmarshalECDSASignature(certFromFile.Signature)
	assert.NoError(t, err)
	lowS, err := utils.IsLowS(caCertFromFile.PublicKey.(*ecdsa.PublicKey), S)
	assert.NoError(t, err)
	assert.False(t, lowS)

	// Check that id.(*signingidentity).cert is in LoswS
	_, S, err = utils.UnmarshalECDSASignature(id.(*signingidentity).cert.Signature)
	assert.NoError(t, err)
	lowS, err = utils.IsLowS(caCertFromFile.PublicKey.(*ecdsa.PublicKey), S)
	assert.NoError(t, err)
	assert.True(t, lowS)

	// Compute the digest for certFromFile
	thisBCCSPMsp := thisMSP.(*bccspmsp)
	hashOpt, err := bccsp.GetHashOpt(thisBCCSPMsp.cryptoConfig.IdentityIdentifierHashFunction)
	assert.NoError(t, err)
	digest, err := thisBCCSPMsp.bccsp.Hash(certFromFile.Raw, hashOpt)
	assert.NoError(t, err)

	// Compare with the digest computed from the sanitised cert
	assert.NotEqual(t, idid.Id, hex.EncodeToString(digest))
}
*/

func TestCLAnonymityIdentity(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPIdentityAnonymity{AnonymityType: msp.MSPIdentityAnonymity_NOMINAL})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ANONYMITY,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.NoError(t, err)
}

func TestCLAnonymityIdentityPreV12Fail(t *testing.T) {
	id, err := localMspCLV11.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPIdentityAnonymity{AnonymityType: msp.MSPIdentityAnonymity_NOMINAL})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ANONYMITY,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLAnonymityIdentityFail(t *testing.T) {
	id, err := localMspCL.GetDefaultSigningIdentity()
	assert.NoError(t, err)

	principalBytes, err := proto.Marshal(&msp.MSPIdentityAnonymity{AnonymityType: msp.MSPIdentityAnonymity_ANONYMOUS})
	assert.NoError(t, err)

	principal := &msp.MSPPrincipal{
		PrincipalClassification: msp.MSPPrincipal_ANONYMITY,
		Principal:               principalBytes}

	err = id.SatisfiesPrincipal(principal)
	assert.Error(t, err)
}

func TestCLProviderTypeToString(t *testing.T) {
	// Check that the provider type is found for FABRIC
	pt := ProviderTypeToString(IBPCLA)
	assert.Equal(t, "ibpcla", pt)
}

func getCLIdentity(t *testing.T, path string) Identity {
	mspDir, err := configtest.GetDevCLMspDir()
	assert.NoError(t, err)

	pems, err := getPemMaterialFromDir(filepath.Join(mspDir, path))
	assert.NoError(t, err)

	id, err := localMspCL.(*clmsp).getclIdentityFromConf(pems[0])
	assert.NoError(t, err)

	return id
}
