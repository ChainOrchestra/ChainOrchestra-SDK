#!/bin/bash 

nodeType=$(echo `basename $0` | sed "s/-.*$//")

if [ $# -lt 3 ]
then
    echo "Usage: `basename $0` host binary archive "
    exit 1
fi

[ "$1" ] && HOST="$1"
[ "$2" ] && BIN="$2"
[ "$3" ] && ARCHIVE="$3"

# Create the bin folder
ssh $HOST "mkdir -p /opt/hyperledger/${nodeType}/bin"

# copy the binary on the remote host
cp $BIN .
tar cz `basename $BIN` | ssh $HOST "tar xz -C /opt/hyperledger/${nodeType}/bin"
rm `basename $BIN`

# Un-tar the archive config files
cat $ARCHIVE | ssh $HOST "tar xz -C /opt/hyperledger/${nodeType}"

# check if everything went well
ssh $HOST 'find /opt/hyperledger'
