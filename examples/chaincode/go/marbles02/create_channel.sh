. ./setenv.sh

peer logging setlevel gossip DEBUG

configtxgen -channelID mycp -outputCreateChannelTx mycp.tx -profile SampleSingleMSPChannel

peer channel create -c mycp -o 127.0.0.1:7050 -f mycp.tx

peer channel join -b mycp.block
