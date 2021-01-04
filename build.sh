#!/bin/sh

rm -rf API
git clone git@github.com:ScienceObjectsDB/API.git

mkdir build

# Builds the go stubs from the protobuf files and places them in their directories
# Temporary solution until ci/cd is enabled 
cd API
git checkout dev
protoc --proto_path=. --go-grpc_out=../build --go_out=../build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/models/*.proto

cd ..
mv build/api/models/CommonModels.pb.go models
mv build/api/models/ObjectModels.pb.go models


rm -rf API
rm -rf build