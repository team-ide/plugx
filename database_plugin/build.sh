#!/bin/bash

CURRENT_DIR=$(pwd)
echo "当前目录：$CURRENT_DIR"

# 获取脚本所在的目录
SCRIPT_DIR="$(dirname "$0")"

# 切换到脚本所在的目录
cd "$SCRIPT_DIR" || exit

SCRIPT_DIR=$(pwd)

echo "脚本目录：$SCRIPT_DIR"

echo "打包 $GOOS $GOARCH dm 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_dm.so common.go dm.go

echo "打包 $GOOS $GOARCH gbase 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_gbase.so common.go gbase.go

echo "打包 $GOOS $GOARCH kingbase_v8r3 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_kingbase_v8r3.so common.go kingbase_v8r3.go

echo "打包 $GOOS $GOARCH kingbase_v8r6 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_kingbase_v8r6.so common.go kingbase_v8r6.go

echo "打包 $GOOS $GOARCH mysql 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_mysql.so common.go mysql.go

echo "打包 $GOOS $GOARCH odbc 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_odbc.so common.go odbc.go

echo "打包 $GOOS $GOARCH opengauss 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_opengauss.so common.go opengauss.go

echo "打包 $GOOS $GOARCH oracle 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_oracle.so common.go oracle.go

echo "打包 $GOOS $GOARCH postgresql 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_postgresql.so common.go postgresql.go

echo "打包 $GOOS $GOARCH shentong 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_shentong.so common.go shentong.go

echo "打包 $GOOS $GOARCH sqlite 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_sqlite.so common.go sqlite.go

echo "打包 $GOOS $GOARCH sqlite_noc 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_sqlite_noc.so common.go sqlite_noc.go

echo "打包 $GOOS $GOARCH ux 插件"
go build -buildmode=plugin -ldflags="-s -w" -o ../plugins/database_ux.so common.go ux.go
