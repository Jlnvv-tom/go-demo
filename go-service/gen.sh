#!/usr/bin/env bash

# 服务名
service=$1

# 类型 api、proto
type=$2

if [[ "$type" == "api" ]]; then
    echo "gen api start"
    goctl api format --dir app/${service}/api/desc
    goctl api go -api app/${service}/api/desc/${service}.api -dir app/${service}/api/  --style=goZero
    elif [[ "$type" = "proto" ]]; then
        echo "gen rpc start"
        echo "goctl rpc protoc proto/${service}/${service}.proto --go_out=proto/${service}/ --go-grpc_out=proto/${service}/  --zrpc_out=app/${service}/rpc/ --style=goZero" -m=true
        goctl rpc protoc proto/${service}/${service}.proto --go_out=proto/${service}/ --go-grpc_out=proto/${service}/  --zrpc_out=app/${service}/rpc/ --style=goZero -m=true

fi

go work init
go work use -r app/*  
cd app/${service} && go mod tidy

#exec /bin/bash