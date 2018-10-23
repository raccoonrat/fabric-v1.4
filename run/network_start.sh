#!/bin/bash
#
#Fabric V1.2 native dev-env 

#1.Please make sure your folder name is exactly "run", then place it in $GOPATH/github.com/hyperledger/fabric/

#2.prepare $GOPATH/github.com/hyperledger/fabric/examples/chaincode/go/marbles02_private (according to the chaincode)

#Follow these scripts for testing ths env. (the "2" in script's name means that it will be executed in the environment of the second peer)
#


bash create_channel.sh
sleep 3
bash join_channel.sh
sleep 3
bash join_channel2.sh
sleep 3
bash set_anchor.sh
sleep 3
bash set_anchor2.sh
sleep 3

bash install_chaincode.sh 
sleep 3
bash install_chaincode2.sh
sleep 3
bash instantiate.sh
sleep 3
#bash invoke.sh


#bash query.sh
#sleep 2
#bash query2.sh

#sleep 2
#bash transfer_marble2.sh
#sleep 2
#bash query.sh
#sleep 2
#bash query2.sh



