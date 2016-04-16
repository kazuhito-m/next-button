#!/bin/bash

source ~/.bashrc

# 依存性解決のgo get(ソースから自動解決)
go get
# テスト
go test ./...
if [[ $? -ne 0 ]] ; then
  echo 'Test failed and Build failed!'
  exit 1
fi
# ビルド・インストール
go install
