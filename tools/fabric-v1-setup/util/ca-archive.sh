#!/bin/bash

nodeType=$(echo `basename $0` | sed "s/-.*$//")

if [ $# -lt 3 ]
then
    echo "Usage: `basename $0` org ${nodeType} template <crypto-config>"
    exit 1
fi

[ "$1" ] && ORG_NAME="$1"
[ "$2" ] && CA_NAME="$2"
[ "$3" ] && TEMPLATE="$3"
CRYPTO_CONFIG="$(pwd)/crypto-config"
[ "$4" ] && CRYPTO_CONFIG="$4"

# run ca-crypto.sh with ORG_NAME and CRYPTO_CONFIG
$(dirname $0)/ca-crypto.sh "$ORG_NAME" "$CRYPTO_CONFIG"
if [ ! $? -eq 0 ]
then
	exit 1
fi
CRYPTO_ARCHIVE="crypto_ca.$ORG_NAME.tgz"

# find ca cert and key 
CA_FOLDER="$(find ${CRYPTO_CONFIG} -type d -name ${ORG_NAME})"/ca
CA_KEYFILE=`basename $(ls "$CA_FOLDER"/*_sk)`
CA_CERTFILE=`basename $(ls "$CA_FOLDER"/*.pem)`
TLS_KEYFILE="$CA_KEYFILE"
TLS_CERTFILE="$CA_CERTFILE"

# run ca-template.sh with CA_NAME and the found certs and keys
$(dirname $0)/ca-template.sh "$TEMPLATE" "$CA_NAME" "$CA_CERTFILE" "$CA_KEYFILE" "$TLS_CERTFILE" "$TLS_KEYFILE"
if [ ! $? -eq 0 ]
then
	exit 1
fi

TMPDIR=$(mktemp -d)
CFGDIR=${TMPDIR}/config
mkdir -p ${CFGDIR}

mv ${CRYPTO_ARCHIVE} ${CFGDIR}
mv `basename $TEMPLATE` ${CFGDIR}

pushd ${CFGDIR}
tar -xzf ${CRYPTO_ARCHIVE}
rm ${CRYPTO_ARCHIVE}
cd ..
CONFIG_ARCHIVE=$(echo ${CRYPTO_ARCHIVE} | sed "s/crypto_/config_/")
tar -czf ${CONFIG_ARCHIVE} config
popd
mv ${TMPDIR}/${CONFIG_ARCHIVE} .
rm -rf ${TMPDIR}


