#!/bin/bash 

cd /opt/hyperledger/orderer
export FABRIC_CFG_PATH=$(pwd)/config
screen -S orderer -L -dm ./bin/orderer
screen -ls
