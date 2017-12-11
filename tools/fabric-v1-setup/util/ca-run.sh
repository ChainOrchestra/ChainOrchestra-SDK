#!/bin/bash 

[ "$1" ] && ADMIN="$1"
[ "$2" ] && PASSWORD="$2"

cd /opt/hyperledger/ca/config/
sed -i -e "s/name: admin/name: $ADMIN/" -e "s/pass: adminpw/pass: $PASSWORD/" fabric-ca-server-config.yaml 
screen -S ca -dm ../bin/fabric-ca-server start
sleep 10
rm fabric-ca-server-config.yaml 
screen -ls
