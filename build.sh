#!/bin/sh

rm -rf API
git clone git@github.com:ScienceObjectsDB/API.git

mkdir build
mkdir -p build/api/openapi

# Builds the go stubs from the protobuf files and places them in their directories
# Temporary solution until ci/cd is enabled 
cd API
git checkout dev
protoc --proto_path=. --go-grpc_out=../build --go_out=../build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/models/*.proto
protoc --proto_path=. --go-grpc_out=../build --go_out=../build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/services/*.proto

protoc -I. --grpc-gateway_out=logtostderr=true,allow_delete_body=true,paths=source_relative:../build api/services/*Service.proto api/services/ObjectLoad.proto

protoc -I . --openapiv2_out ../build/api/openapi --openapiv2_opt logtostderr=true,allow_merge=true api/services/*Service.proto api/services/ObjectLoad.proto

cd ..
mv build/api/models/* models
mv build/api/services/* services
mv -f build/api/openapi/* openapi

go generate openapi/SwaggerGen.go

rm -rf API
rm -rf build