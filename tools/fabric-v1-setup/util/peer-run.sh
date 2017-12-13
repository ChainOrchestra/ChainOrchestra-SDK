#!/bin/bash 

cd /opt/hyperledger/peer
export FABRIC_CFG_PATH=$(pwd)/config
screen -S peer -L -dm ./bin/peer node start
screen -ls
