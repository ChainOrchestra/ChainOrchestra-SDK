## Initial blockchain setup

```
FABRIC_PATH=/opt/gopath/src/github.com/hyperledger/fabric/build/bin
export PATH=$FABRIC_PATH:$PATH

cd sample

# generate crypto-config
cryptogen generate --config=crypto-config.yaml --output=crypto-config

# generate orderer genesis block
configtxgen -profile OrdererGenesis -outputBlock orderer.genesis.block

# generate Org1 channel configuration transaction
configtxgen -profile Org1Channel -outputCreateChannelTx org1.tx -channelID org1chan

# generate Org2 channel configuration transaction
configtxgen -profile Org2Channel -outputCreateChannelTx org2.tx -channelID org2chan

# generate shared channel configuration transaction
configtxgen -profile SharedChannel -outputCreateChannelTx shared.tx -channelID sharedchan
```

