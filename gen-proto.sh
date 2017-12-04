#!/usr/bin/env bash

LG='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

for f in service/*/
do
    echo -e "${RED}[INFO]${NC} Processing proto buffer for ${LG}$(basename $f) ${NC}service${NC}"
    protoc -I=. -I=$GOPATH/src --go_out=plugins=grpc:. $f/proto/*/*.proto
done