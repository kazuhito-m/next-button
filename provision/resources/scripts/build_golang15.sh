#!/bin/bash

# go1.5のダウンロード＆ビルド＆インストール
# http://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi を参考に

WORK_DIR="/tmp/golang"

rm -rf $WORK_DIR
mkdir $WORK_DIR
cd $WORK_DIR

curl http://dave.cheney.net/paste/go-linux-arm-bootstrap-c788a8e.tbz | tar xj
curl https://storage.googleapis.com/golang/go1.5.1.src.tar.gz | tar xz

ulimit -s 1024    #32bitなので
cd $WORK_DIR/go/src
env GO_TEST_TIMEOUT_SCALE=10 GOROOT_BOOTSTRAP=$HOME/go-linux-arm-bootstrap ./all.bash

rm -rf $WORK_DIR/go-linux-arm-bootstrap    #不要なので削除
