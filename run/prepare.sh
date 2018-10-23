#!/bin/bash
#
# Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

rm -rf crypto-config channel-artifacts
mkdir channel-artifacts
set FABRIC_CFG_PATH=`pwd`
set CHANNEL_NAME=mychannel
cp ../.build/bin/* .

./cryptogen generate --config=./crypto-config.yaml

IDEMIXMATDIR=$CURDIR/crypto-config/idemix
mkdir -p ./crypto-config/idemix
./idemixgen ca-keygen
./idemixgen signerconfig -u OU1 -e OU1 -r 1

cp -r ./idemix-config ./crypto-config/idemix/

rm -rf channel-artifacts
mkdir -p channel-artifacts
./configtxgen -profile TwoOrgsOrdererGenesis -channelID e2e-orderer-syschan -outputBlock ./channel-artifacts/genesis.block

#create the channel transaction artifact
export CHANNEL_NAME=mychannel  && ./configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME

#define the anchor peer for Org1 on the channel that we are constructing.
./configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

#define the anchor peer for Org2 on the same channel:
./configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
