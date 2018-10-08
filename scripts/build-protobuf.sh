#!/usr/bin/env bash

LATEST_VEIDEMANN_API_URL=$(curl -Ls -o /dev/null -w %{url_effective} https://github.com/nlnwa/veidemann-api/releases/latest)
API_VERSION=${LATEST_VEIDEMANN_API_URL##*/}

LATEST_PROTOBUF_URL=$(curl -Ls -o /dev/null -w %{url_effective} https://github.com/protocolbuffers/protobuf/releases/latest)
PROTOBUF_VERSION=${LATEST_PROTOBUF_URL##*/v}

echo -e "Veidemann API version:\t${API_VERSION}"
echo -e "Protobuf version:\t${PROTOBUF_VERSION}"

rm -rf protobuf veidemann_api
mkdir protobuf veidemann_api
cd protobuf

wget -q https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
unzip protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
rm protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
wget -O - -q https://github.com/nlnwa/veidemann-api/archive/${API_VERSION}.tar.gz | tar --strip-components=2 -zx
go get -u github.com/golang/protobuf/protoc-gen-go

bin/protoc -I. --go_out=plugins=grpc:../veidemann_api *.proto