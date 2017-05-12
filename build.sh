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

echo "Building ./api"
pushd $dir/api > /dev/null
make build
popd > /dev/null