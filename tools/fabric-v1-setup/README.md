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


