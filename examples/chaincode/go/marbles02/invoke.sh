. ./setenv.sh

peer chaincode invoke -o 127.0.0.1:7050 -C mycp -n marblesp  -c '{"Args":["initMarble","marble1","blue","35","tom"]}'

#peer chaincode invoke -o 127.0.0.1:7050 -C mycp -n marbles  -c '{"Args":["initMarble","marble2","red","50","tom"]}'

