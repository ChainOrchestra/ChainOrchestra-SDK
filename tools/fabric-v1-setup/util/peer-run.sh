#!/bin/bash 

cd /opt/hyperledger/peer
export FABRIC_CFG_PATH=$(pwd)/config
screen -S peer -dm ./bin/peer node start
screen -ls
