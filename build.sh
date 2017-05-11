#!/bin/bash

dir=`pwd`

pushd $dir/srv/joke > /dev/null
echo "Building ./srv/joke"
make build
popd > /dev/null


pushd $dir/srv/name > /dev/null
echo "Building ./srv/name"
make build
popd > /dev/null

# echo "Building ./api"
# pushd $dir/api > /dev/null
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w'
# popd > /dev/null