rem # 1. 
rm -rf crypto-config channel-artifacts
mkdir channel-artifacts
set FABRIC_CFG_PATH=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run
set CHANNEL_NAME=mychannel


cryptogen generate --config=./crypto-config.yaml


rem #2.Create the orderer genesis block:
    configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

rem #3. define the channel transaction artifact:
   configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID %CHANNEL_NAME%


rem #4.Define the anchor peer for Org1 on the channel:
    configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID %CHANNEL_NAME% -asOrg Org1MSP


rem #5.Define the anchor peer for Org2 on the channel:
    configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID %CHANNEL_NAME% -asOrg Org2MSP

