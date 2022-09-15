#! /bin/bash

# Run :daemon
# ./bisq-daemon --apiPassword=xyz  --appDataDir=/tmp/newbisqdatadir
# Run go grpc proxy
# go run main.go proxy

x=1
while [ $x -le 100 ]
do
	echo "proxy test iteration $x"
	
	./curl-cmds.sh

	x=$(( $x + 1 ))

	sleep 10
	
done
 
