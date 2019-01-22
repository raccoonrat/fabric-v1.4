
set PROJECT_VERSION=1.0.5
set BASEIMAGE_RELEASE=0.4.6
set CORE_LOGGING_LEVEL=DEBUG
set FABRIC_CFG_PATH=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run
set FABRIC_ROOT=C:/Workspace/gopath/src/github.com/hyperledger/fabric
set ORDERER_GENERAL_LOGLEVEL=DEBUG
set ORDERER_GENERAL_TLS_ENABLED=false
set ORDERER_GENERAL_PROFILE_ENABLED=false
set ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
set ORDERER_GENERAL_LISTENPORT=7050
set ORDERER_GENERAL_GENESISMETHOD=file
set ORDERER_GENERAL_GENESISFILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/channel-artifacts/genesis.block
set ORDERER_GENERAL_LOCALMSPDIR=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp
set ORDERER_GENERAL_LOCALMSPID=OrdererMSP
set ORDERER_FILELEDGER_LOCATION=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/hyperledger/product/orderer
set CORE_ORDERER_TLS_CERT_FILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderers.example.com/tls/server.crt
set CORE_ORDERER_TLS_KEY_FILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderers.example.com/tls/server.key
set CORE_ORDERER_TLS_ROOTCERT_FILE=C:/Workspace/gopath/src/github.com/hyperledger/fabric/run/crypto-config/ordererOrganizations/example.com/orderers/orderers.example.com/tls/ca.crt



orderer start
