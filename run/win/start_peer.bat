set FABRIC_CFG_PATH=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run
set CORE_LOGGING_LEVEL=DEBUG
set CORE_PEER_TLS_ENABLED=false
set CORE_PEER_GOSSIP_USELEADERELECTION=true
set CORE_PEER_GOSSIP_ORGLEADER=false
set CORE_PEER_PROFILE_ENABLED=true
set CORE_PEER_MSPCONFIGPATH=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
set CORE_PEER_ID=peer0.org1.example.com
set CORE_PEER_ADDRESS=peer0.org1.example.com:7051
rem set CORE_PEER_CHAINCODELISTENADDRESS=10.116.22.103:7052
set CORE_PEER_CHAINCODELISTENADDRESS=192.168.1.106:7052
set CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org1.example.com:7051
set CORE_PEER_TLS_CERT_FILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt
set CORE_PEER_TLS_KEY_FILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key
set CORE_PEER_TLS_ROOTCERT_FILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
set GOPATH=C:/Workspace/gopath
set CORE_PEER_LOCALMSPID=Org1MSP

peer node start  -o 127.0.0.1:7050
