#!/bin/bash

#
# next-button のデーモンスクリプト。
#
# 無限ループにて「ボタンが押されたかどうか」を判定する。

# 自身のスクリプトのパスを取得する。
CUR_DIR=$(cd $(dirname $0);pwd)
LIB_DIR=/usr/local/next-button
STATUS_FILE=/tmp/next-button.status
LOG_FILE=/var/log/next-button.log
BUTTUN_EVENT=${LIB_DIR}/next-button.bsh

# GPIO系定数
GPIO_NO=2
INTERVAL_SEC=0.1
GPIO_DIR=/sys/class/gpio
GPIO_VALUE=${GPIO_DIR}/gpio${GPIO_NO}/value

# GPIO周りを初期化
echo ${GPIO_NO} > ${GPIO_DIR}/export
sleep 1
echo in > ${GPIO_DIR}/gpio${GPIO_NO}/direction

# 初回ステータスファイルを入れて置く
echo '0' > ${STATUS_FILE}

while true ; do
  # インターバルスリープ
  sleep ${INTERVAL_SEC}
  # ボタンのステータスを取得
  now_button_status=`cat ${GPIO_VALUE}`
  # 変化していたら
  if [[ ${last_button_status} != ${now_button_status} ]] ; then
    echo ${now_button_status} > ${STATUS_FILE}
    last_button_status=${now_button_status}
    # 0->1へ変化した時のみ
    if [[ ${now_button_status} = 1 ]] ; then
      date >> ${LOG_FILE}
      echo "異なった場合のルーチンに入ってきている : ${now_button_status}" >>  ${LOG_FILE}
      # 「実際に行いたい処理」のスクリプトを
      ${BUTTUN_EVENT}
    fi
  fi
done

# コレが実行されることはない…うまいことしたい
echo ${GPIO_NO} > /sys/class/gpio/unexport

exit 0
