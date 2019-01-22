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
rm -rf mycp.*
