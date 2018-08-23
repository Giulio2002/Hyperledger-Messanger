docker exec cli0.Org1 bash -c "peer chaincode invoke -C org1 -n Authenticator -c '{\"Args\":[\"register\", \"FunnyGiulio\",\"Password\"]}' --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/messanger.com/orderers/orderer.messanger.com/msp/tlscacerts/tlsca.messanger.com-cert.pem"
docker exec cli0.Org1 bash -c "peer chaincode invoke -C org1 -n Authenticator -c '{\"Args\":[\"register\", \"FunnyLuigi\",\"Password\"]}' --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/messanger.com/orderers/orderer.messanger.com/msp/tlscacerts/tlsca.messanger.com-cert.pem"
sleep 1
docker exec cli0.Org1 bash -c "peer chaincode invoke -C org1 -n Authenticator -c '{\"Args\":[\"exist\", \"FunnyLuigi\"]}' --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/messanger.com/orderers/orderer.messanger.com/msp/tlscacerts/tlsca.messanger.com-cert.pem"