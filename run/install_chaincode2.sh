#!/bin/bash
#
# Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

export  FABRIC_CFG_PATH=$GOPATH/src/github.com/hyperledger/fabric/run
export  CORE_LOGGING_LEVEL=DEBUG
export CORE_CHAINCODE_LOGGING_LEVEL=DEBUG
export  CORE_PEER_TLS_ENABLED=false
export  CORE_PEER_GOSSIP_USELEADERELECTION=true
export  CORE_PEER_GOSSIP_ORGLEADER=false
export  CORE_PEER_PROFILE_ENABLED=true
export  CORE_PEER_TLS_CERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.crt
export  CORE_PEER_TLS_KEY_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.key
export  CORE_PEER_TLS_ROOTCERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderers.example.com/tls/ca.crt

export  CORE_PEER_MSPCONFIGPATH=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp

#added for second peer
export CORE_PEER_LISTENADDRESS=127.0.0.1:7055
export CORE_PEER_GOSSIP_BOOTSTRAP=127.0.0.1:7055
export CORE_PEER_EVENTS_ADDRESS=127.0.0.1:7058
export CORE_PEER_PROFILE_LISTENADDRESS=127.0.0.1:9090

export CORE_PEER_FILESYSTEMPATH=./hyperledger/production2
export  CORE_PEER_ID=cli
export  CORE_PEER_ADDRESS=127.0.0.1:7055
export  CORE_PEER_CHAINCODELISTENADDRESS=127.0.0.1:7056
export  CORE_PEER_GOSSIP_EXTERNALENDPOINT=127.0.0.1:7055
export  CORE_PEER_LOCALMSPID=Org2MSP
export  CHANNEL_NAME=mychannel

#./peer chaincode install -n mycc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/map

./peer chaincode install -n marblesp -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/marbles02_private/go/


#docker ps -a

#./peer chaincode instantiate -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -v 1.0 -c "{\"Args\":[\"init\",\"a\", \"100\", \"b\",\"200\"]}" -P "OR('OrgMSP.member')"
#./peer chaincode instantiate -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -v 1.0  -c '{"Args":["invoke","put","b","10"]}'

#./peer chaincode invoke -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -c '{"Args":["invoke","a","b","10"]}'
