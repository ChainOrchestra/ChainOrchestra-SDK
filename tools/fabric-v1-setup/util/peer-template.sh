#!/bin/bash

nodeType=$(echo `basename $0` | sed "s/-.*$//")
if [ "$nodeType" = "peer" ]
then
	subst=( "PEERID" "NETWORKID" "ADDRESS" "BOOTSTRAP" "LOCALMSPID" )
elif [ "$nodeType" = "orderer" ]
then
	subst=( "LOCALMSPID" )
elif [ "$nodeType" = "ca" ]
then
	subst=( "CA_NAME" "CA_CERTFILE" "CA_KEYFILE" "TLS_CERTFILE" "TLS_KEYFILE" )
fi


if [ $# -ne $(( 1 + ${#subst[@]})) ]
then
    echo "Usage: `basename $0` template ${subst[@]}"
    exit 1
fi

TEMPLATE_FILE="$1"

if [ ! -f ${TEMPLATE_FILE} ]
then
	echo "Cannot find template file ${TEMPLATE_FILE}"
	exit 1
fi

TARGET_FILE=$(basename ${TEMPLATE_FILE})
cp ${TEMPLATE_FILE} ${TARGET_FILE}

for i in ${subst[@]}
do
	shift
	sed -i "s/$i/$1/g" ${TARGET_FILE}
done
