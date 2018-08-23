export FABRIC_CFG_PATH=./

configtxgen -profile messangerOrdererGenesis -outputBlock ./orderer/genesis.block
configtxgen -profile Org1Channel -outputCreateChannelTx ./channels/Org1.tx -channelID org1
configtxgen -profile Org2Channel -outputCreateChannelTx ./channels/Org2.tx -channelID org2
configtxgen -profile Org1Channel -outputAnchorPeersUpdate ./channels/org1anchor.tx -channelID org1 -asOrg Org1MSP
configtxgen -profile Org2Channel -outputAnchorPeersUpdate ./channels/org2anchor.tx -channelID org2 -asOrg Org2MSP
configtxgen -profile Org1Channel -outputAnchorPeersUpdate ./channels/Authanchororg1.tx -channelID org1 -asOrg AuthMSP
configtxgen -profile Org2Channel -outputAnchorPeersUpdate ./channels/Authanchororg2.tx -channelID org2 -asOrg AuthMSP
