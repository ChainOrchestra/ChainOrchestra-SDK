#!/bin/bash

nodeType=$(echo `basename $0` | sed "s/-.*$//")

if [ $# -lt 2 ]
then
    echo "Usage: `basename $0` org ${nodeType} <crypto-config>"
    exit 1
fi

[ "$1" ] && ORG_NAME="$1"
[ "$2" ] && PEER_NAME="$2"
CRYPTO_CONFIG="$(pwd)/crypto-config"
[ "$3" ] && CRYPTO_CONFIG="$3"

if [ ! -d ${CRYPTO_CONFIG}/${nodeType}Organizations/${ORG_NAME}/${nodeType}s/${PEER_NAME} ]
then
	echo "Cannot find folder ${CRYPTO_CONFIG}/${nodeType}Organizations/${ORG_NAME}/${nodeType}s/${PEER_NAME}"
	exit 1
fi

cd ${CRYPTO_CONFIG}/${nodeType}Organizations/${ORG_NAME}/${nodeType}s/${PEER_NAME}
tar -czf crypto_${PEER_NAME}.tgz msp tls
cd -
mv ${CRYPTO_CONFIG}/${nodeType}Organizations/${ORG_NAME}/${nodeType}s/${PEER_NAME}/crypto_${PEER_NAME}.tgz .
