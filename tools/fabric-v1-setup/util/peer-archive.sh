#!/bin/bash

nodeType=$(echo `basename $0` | sed "s/-.*$//")
cryptoSh=$(echo $0 | sed "s/-archive\.sh$/-crypto\.sh/")

if [ $# -lt 4 ]
then
    echo "Usage: `basename $0` org ${nodeType} crypto-config template [SUBSTITUTIONS]"
    exit 1
fi

[ "$1" ] && ORG_NAME="$1"
[ "$2" ] && PEER_NAME="$2"
[ "$3" ] && CRYPTO_CONFIG="$3"

# run xxx-crypto.sh with 3 first parameters
${cryptoSh} "$ORG_NAME" "$PEER_NAME" "$CRYPTO_CONFIG"
if [ ! $? -eq 0 ]
then
	exit 1
fi
CRYPTO_ARCHIVE="crypto_${PEER_NAME}.tgz"

shift; shift; shift

# in the case of an orderer, run orderer-template.sh with orderer.yaml template and LOCALMSPID
if [ "$nodeType" = "orderer" ]
then
	$(dirname $0)/orderer-template.sh $(echo $1 | sed "s/core/orderer/") "$6"
	if [ ! $? -eq 0 ]
	then
		exit 1
	fi
	
fi

# run peer-template.sh with the remaining parameters
$(dirname $0)/peer-template.sh $*
if [ ! $? -eq 0 ]
then
	exit 1
fi

TMPDIR=$(mktemp -d)
CFGDIR=${TMPDIR}/config
mkdir -p ${CFGDIR}

# in the case of an orderer, also install orderer.yaml from template folder
# and install orderer.genesis.block from crypto-config parent folder
mv ${CRYPTO_ARCHIVE} ${CFGDIR}
if [ "$nodeType" = "orderer" ]
then
	mv orderer.yaml ${CFGDIR}
	cp ${CRYPTO_CONFIG}/../orderer.genesis.block ${CFGDIR}
fi
mv core.yaml ${CFGDIR}

pushd ${CFGDIR}
tar -xzf ${CRYPTO_ARCHIVE}
rm ${CRYPTO_ARCHIVE}
cd ..
CONFIG_ARCHIVE=$(echo ${CRYPTO_ARCHIVE} | sed "s/crypto_/config_/")
tar -czf ${CONFIG_ARCHIVE} config
popd
mv ${TMPDIR}/${CONFIG_ARCHIVE} .
rm -rf ${TMPDIR}


