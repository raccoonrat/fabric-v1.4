rm -rf channel-artifacts  crypto-config      hyperledger 
rm -rf ./mychannel.block
rm -rf chaintool  configtxgen configtxlator  cryptogen orderer  peer block-listener
docker rm -f $(docker ps -aq)

orderer=""
peer=""
orderer=`ps -ef|grep orderer |grep -v grep | awk '{print $2}'`
peer=`ps -ef|grep peer|grep -v grep | awk '{print $2}'`
if [ "x$orderer" != "x" ]
then
#kill $orderer
echo $orderer
fi

if [ "x$peer" != "x" ]
then
echo $orderer
#kill $peer
fi

rm -rf /var/hyperledger
