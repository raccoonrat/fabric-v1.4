. ./setenv.sh

#peer chaincode install -n marblesp -p github.com/hyperledger/fabric/samples/chaincode/marbles02/go -v 1
peer chaincode install -n marblesp -p github.com/hyperledger/fabric/examples/chaincode/go/marbles02 -v 1

#peer chaincode instantiate -C mycp -n marblesp -c '{"Args":["init"]}' -v 1 -o 127.0.0.1:7050 --collections-config collection.json
