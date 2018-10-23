export  FABRIC_CFG_PATH=$GOPATH/src/github.com/hyperledger/fabric/run
export  CORE_LOGGING_LEVEL=DEBUG
export  CORE_PEER_TLS_ENABLED=false
export  CORE_PEER_GOSSIP_USELEADERELECTION=true
export  CORE_PEER_GOSSIP_ORGLEADER=false
export  CORE_PEER_PROFILE_ENABLED=true
export  CORE_PEER_TLS_CERT_FILE=/home/wangyh/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt
export  CORE_PEER_TLS_KEY_FILE=/home/wangyh/gopath/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key
#export  CORE_PEER_TLS_ROOTCERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderers.example.com/tls/ca.crt
export  CORE_PEER_TLS_ROOTCERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

export  CORE_PEER_MSPCONFIGPATH=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

export  CORE_PEER_ID=cli
export  CORE_PEER_ADDRESS=127.0.0.1:7051
export  CORE_PEER_CHAINCODELISTENADDRESS=127.0.0.1:7052
export  CORE_PEER_GOSSIP_EXTERNALENDPOINT=127.0.0.1:7051
export  CORE_PEER_LOCALMSPID=Org1MSP
export  CHANNEL_NAME=mychannel



./peer channel create -o 127.0.0.1:7050 -c mychannel -f ./channel-artifacts/channel.tx 




