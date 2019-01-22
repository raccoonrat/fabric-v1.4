. ./setenv.sh

peer chaincode query -C mycp -n marblesp -c '{"Args":["readMarble","marble1"]}'

peer chaincode query -C mycp -n marblesp -c '{"Args":["readMarblePrivateDetails","marble1"]}'

peer chaincode query -C mycp -n marblesp -c '{"Args":["getMarblesByRange","marble1","marble8"]}'
