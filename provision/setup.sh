#!/bin/bash

# RaspberryPI に Raspabianおインストール後、
# Lanケーブルをつなぎ、DHCPで取得するIPを調べた時点で、
# 以下を実行
fab -H 192.168.1.90 -u pi -p raspberry setup_all 

