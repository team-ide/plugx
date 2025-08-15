#!/bin/bash

CURRENT_DIR=$(pwd)
echo "当前目录：$CURRENT_DIR"

# 获取脚本所在的目录
SCRIPT_DIR="$(dirname "$0")"

# 切换到脚本所在的目录
cd "$SCRIPT_DIR" || exit

SCRIPT_DIR=$(pwd)

echo "脚本目录：$SCRIPT_DIR"

echo "打包 $GOOS $GOARCH sqlite 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../../plugins/dbx_sqlite.so common.go sqlite.go

echo "打包 $GOOS $GOARCH mysql 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../../plugins/dbx_mysql.so common.go mysql.go

echo "打包 $GOOS $GOARCH sqlite_noc 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../../plugins/dbx_sqlite_noc.so common.go sqlite_noc.go
