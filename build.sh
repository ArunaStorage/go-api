#!/bin/sh

rm -rf API
git clone git@github.com:ScienceObjectsDB/API.git

mkdir build

# Builds the go stubs from the protobuf files and places them in their directories
# Temporary solution until ci/cd is enabled 
cd API
git checkout dev
protoc --proto_path=. --go-grpc_out=../build --go_out=../build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/models/*.proto
protoc --proto_path=. --go-grpc_out=../build --go_out=../build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/services/*.proto


cd ..
mv build/api/models/* models
mv build/api/services/* services

rm -rf API
rm -rf build