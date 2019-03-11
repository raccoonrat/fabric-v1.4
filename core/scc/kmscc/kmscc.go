package kmscc

import (
	"fmt"

	"github.com/hyperledger/fabric/core/aclmgmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/core/common/sysccprovider"
)

func New(ccp ccprovider.ChaincodeProvider,
	sccp sysccprovider.SystemChaincodeProvider,
	aclProvider aclmgmt.ACLProvider) *KMSCC {

	return &KMSCC{}
}

// KMSCC the Key Management System Chain Code
type KMSCC struct {
}

func (km *KMSCC) Name() string {
	return "kmscc"
}

func (km *KMSCC) Path() string {
	return "github.com/hyperledger/fabric/core/scc/kmscc"
}

func (km *KMSCC) InitArgs() [][]byte {
	return nil
}

func (km *KMSCC) Chaincode() shim.Chaincode {
	return km
}

func (km *KMSCC) InvokableExternal() bool {
	return true
}

func (km *KMSCC) InvokableCC2CC() bool {
	return true
}

func (km *KMSCC) Enabled() bool {
	return true
}

func (km *KMSCC) Init(stub shim.ChaincodeStubInterface) shim.Response {
	return shim.Success([]byte("OK"))
}

func (km *KMSCC) Invoke(stub shim.ChaincodeStubInterface) shim.Response {
	args := stub.GetArgs()

	if len(args) < 1 {
		return shim.Error(fmt.Sprintf("Incorrect number of arguments, %d", len(args)))
	}

	fname := string(args[0])

	// if fname != GetChannels && len(args) < 2 {
	// 	return shim.Error(fmt.Sprintf("Incorrect number of arguments, %d", len(args)))
	// }

	//cnflogger.Debugf("Invoke function: %s", fname)

	// Handle ACL:
	// 1. get the signed proposal
	sp, err := stub.GetSignedProposal()
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed getting signed proposal from stub: [%s]", err))
	}

	return shim.Success([]byte("OK"))
}
