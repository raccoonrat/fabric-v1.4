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
export  CORE_PEER_TLS_CERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt
export  CORE_PEER_TLS_KEY_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key
export  CORE_PEER_TLS_ROOTCERT_FILE=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderers.example.com/tls/ca.crt

export  CORE_PEER_MSPCONFIGPATH=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

export  CORE_PEER_ID=cli
export  CORE_PEER_ADDRESS=127.0.0.1:7051
export  CORE_PEER_CHAINCODELISTENADDRESS=127.0.0.1:7052
export  CORE_PEER_GOSSIP_EXTERNALENDPOINT=127.0.0.1:7051
export  CORE_PEER_LOCALMSPID=Org1MSP
export  CHANNEL_NAME=mychannel
export  ORDERER_CA=$GOPATH/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem


#docker ps -a

#./peer chaincode instantiate -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -v 1.0 -c "{\"Args\":[\"init\",\"a\", \"100\", \"b\",\"200\"]}" -P "OR('OrgMSP.member')"
#./peer chaincode instantiate -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -v 1.0  -c '{"Args":["invoke","put","b","10"]}'

#./peer chaincode instantiate -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -v 1.0  -c '{"Args":["init"]}'

./peer chaincode instantiate -o orderer.example.com:7050  -C mychannel -n marblesp -v 1.0 -c '{"Args":["init"]}' -P "OR('Org1MSP.member','Org2MSP.member')" --collections-config  $GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/go/marbles02_private/collections_config.json


#./peer chaincode invoke -o 127.0.0.1:7050 -C $CHANNEL_NAME -n mycc -c '{"Args":["invoke","a","b","10"]}'
