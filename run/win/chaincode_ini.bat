set ABRIC_CFG_PATH=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run
set ORE_LOGGING_LEVEL=DEBUG
set ORE_PEER_LOCALMSPID=Org1MSP
set HANNEL_NAME=mychannel
set ORE_PEER_MSPCONFIGPATH=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
set ORE_PEER_ID=cli
set OPATH=/opt/gopath
set ORE_PEER_ADDRESS=10.116.22.103:7051


peer chaincode install -n mycc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02

peer chaincode instantiate -o 127.0.0.1:7050 -C mychannel -n mycc -v 1.0 -c "{\"Args\":[\"init\",\"a\", \"100\", \"b\",\"200\"]}" -P "OR('OrgMSP.member')"


