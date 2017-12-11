#!/bin/bash

if [ $# -lt 1 ]
then
    echo "Usage: `basename $0` org <crypto-config>"
    exit 1
fi

[ "$1" ] && ORG_NAME="$1"
CRYPTO_CONFIG="$(pwd)/crypto-config"
[ "$2" ] && CRYPTO_CONFIG="$2"

CA_FOLDER="$(find ${CRYPTO_CONFIG} -type d -name ${ORG_NAME})"/ca

if [ ! -d ${CA_FOLDER} ]
then
	echo "Cannot find folder ${CA_FOLDER}"
	exit 1
fi

cd ${CA_FOLDER}
tar -czf crypto_ca.${ORG_NAME}.tgz *
cd -
mv ${CA_FOLDER}/crypto_ca.${ORG_NAME}.tgz .
