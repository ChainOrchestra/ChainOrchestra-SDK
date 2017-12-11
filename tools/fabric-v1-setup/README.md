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
configtxgen -profile Org1Channel -outputCreateChannelTx org1chan.tx -channelID org1chan

# generate Org2 channel configuration transaction
configtxgen -profile Org2Channel -outputCreateChannelTx org2chan.tx -channelID org2chan

# generate shared channel configuration transaction
configtxgen -profile SharedChannel -outputCreateChannelTx sharedchan.tx -channelID sharedchan
```

## Orderer setup

### Archive for orderer

```
../util/orderer-archive.sh \
	orderer.example.org solo.orderer.example.org crypto-config ../template/core.yaml \
	orderer testnet solo.orderer.example.org 127.0.0.1:7051 OrdererMSP
```

### Deploy on orderer
```
../util/orderer-deploy.sh solo.orderer.example.org \
	/opt/gopath/src/github.com/hyperledger/fabric/build/bin/orderer \
	config_solo.orderer.example.org.tgz
```

### Start orderer
```
scp ../util/orderer-run.sh solo.orderer.example.org:
ssh solo.orderer.example.org "./orderer-run.sh; rm orderer-run.sh"
```


## Peer organizations setup

```
ORG=org1
MSP=Org1MSP
```

```
PEER=peer0
```

### Peer archive

The bootstrap peer is ```peer0.${ORG}.example.org:7051```

```
../util/peer-archive.sh \
	${ORG}.example.org ${PEER}.${ORG}.example.org crypto-config ../template/core.yaml \
	${PEER} testnet ${PEER}.${ORG}.example.org peer0.${ORG}.example.org:7051 ${MSP}
```

### Peer deployment

```
../util/peer-deploy.sh ${PEER}.${ORG}.example.org \
	/opt/gopath/src/github.com/hyperledger/fabric/build/bin/peer \
	config_${PEER}.${ORG}.example.org.tgz
```

### Run peer on remote host

```
scp ../util/peer-run.sh ${PEER}.${ORG}.example.org:
ssh ${PEER}.${ORG}.example.org "./peer-run.sh; rm peer-run.sh"
```

## Certificate Authorities management

In addition to the ```$ORG``` and ```$MSP``` variables, we use ```$CA_NAME``` for the Certificate Authority name.
By default it is relative to the organisation name.

```
CA_NAME=ca-${ORG}
```

### Certificate Authority archive

```
../util/ca-archive.sh \
	${ORG}.example.org ${CA_NAME} ../template/fabric-ca-server-config.yaml crypto-config
```

### Certificate Authority deployment

```
../util/ca-deploy.sh ca.${ORG}.example.org \
	/opt/gopath/src/github.com/hyperledger/fabric-ca/bin/fabric-ca-server \
	config_ca.${ORG}.example.org.tgz
```

### Run Certificate Authority server on remote host

  * **Note**: Set the Certificate Authority administrator name and password before launching the script

```
set +o history
ADMIN=myadmin
PASSWORD=mypassword
set -o history
```

```
scp ../util/ca-run.sh ca.${ORG}.example.org:
ssh ca.${ORG}.example.org "./ca-run.sh $ADMIN $PASSWORD; rm ca-run.sh"
```


## Command Line Interface test sessions

### Default channels per organization

```
ORG=org1
MSP=Org1MSP
CHANNEL_NAME=org1chan
```

### Peer environment setup

Choose a peer on the organization selected above.

```
PEER=peer0
```

Set the CRYPTO_CONFIG root for the certification files.

```
CRYPTO_CONFIG=$(pwd)/crypto-config
```

Then set the peer CLI environment.

```
export ORDERER_CA=${CRYPTO_CONFIG}/ordererOrganizations/orderer.example.org/orderers/solo.orderer.example.org/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=${CRYPTO_CONFIG}/peerOrganizations/${ORG}.example.org/peers/${PEER}.${ORG}.example.org/tls/ca.crt
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_MSPCONFIGPATH=${CRYPTO_CONFIG}/peerOrganizations/${ORG}.example.org/users/Admin@${ORG}.example.org/msp
export CORE_PEER_LOCALMSPID=${MSP}
export CORE_PEER_ADDRESS=${PEER}.${ORG}.example.org:7051
```

### Channel management

  * Channel creation. **Save the channel.block** for other peers to join later.

```
peer channel create -c $CHANNEL_NAME -o solo.orderer.example.org:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -f ${CRYPTO_CONFIG}/../$CHANNEL_NAME.tx
```

  * Joining a channel

```
peer channel join -b $CHANNEL_NAME.block
```

### Chaincode management

  * Sample chainconde installation

```
peer chaincode install -n ex02cc -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02
```

  * Sample chaincode instantiation. **Run only on the first peer** where the chaincode is installed.

```
CHANNEL_NAME=org1chan
peer chaincode instantiate -o solo.orderer.example.org:7050 \
	--tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME \
	-n ex02cc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' \
	-P "OR ('Org1MSP.member')"
```

### Test session

  * Sample queries and transactions

```
peer chaincode query -C $CHANNEL_NAME -n ex02cc -c '{"Args":["query","a"]}'
peer chaincode invoke -o solo.orderer.example.org:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n ex02cc -c '{"Args":["invoke","a","b","10"]}'
peer chaincode query -C $CHANNEL_NAME -n ex02cc -c '{"Args":["query","a"]}'
peer chaincode query -C $CHANNEL_NAME -n ex02cc -c '{"Args":["query","b"]}'
```

