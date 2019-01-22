. ./setenv.sh

peer chaincode instantiate -C mycp -n marblesp -c '{"Args":["init"]}' -v 1 -o 127.0.0.1:7050 --collections-config collection.json

