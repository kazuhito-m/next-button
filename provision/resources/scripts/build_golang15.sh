#!/bin/bash

# go1.5のダウンロード＆ビルド＆インストール
# http://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi を参考に

WORK_DIR="/tmp/golang"

rm -rf $WORK_DIR
mkdir $WORK_DIR
cd $WORK_DIR

hg clone -u default https://code.google.com/p/go
cd ./go/src
./make.bash
