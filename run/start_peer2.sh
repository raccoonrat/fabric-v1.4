export FABRIC_CFG_PATH=$GOPATH/src/github.com/hyperledger/fabric/run
export CORE_LOGGING_LEVEL=DEBUG
export CORE_PEER_TLS_ENABLED=false
export CORE_PEER_GOSSIP_USELEADERELECTION=true
export CORE_PEER_GOSSIP_ORGLEADER=false
export CORE_PEER_PROFILE_ENABLED=true
export CORE_PEER_TLS_CERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.crt
export CORE_PEER_TLS_KEY_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.key
export CORE_PEER_TLS_ROOTCERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/msp
export CORE_PEER_ID=peer0.org2.example.com
export CORE_PEER_ADDRESS=127.0.0.1:7055
#added for second peer
export CORE_PEER_LISTENADDRESS=127.0.0.1:7055
export CORE_PEER_GOSSIP_BOOTSTRAP=127.0.0.1:7055
export CORE_PEER_EVENTS_ADDRESS=127.0.0.1:7058
export CORE_PEER_PROFILE_LISTENADDRESS=127.0.0.1:9090
export CORE_PEER_FILESYSTEMPATH=./hyperledger/production2

export CORE_PEER_CHAINCODELISTENADDRESS=127.0.0.1:7056
export CORE_PEER_GOSSIP_EXTERNALENDPOINT=127.0.0.1:7055
export CORE_PEER_LOCALMSPID=Org2MSP

./peer node start 
