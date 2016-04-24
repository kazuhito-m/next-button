package nexttrain

import (
	"testing"
	"time"
)


// 判定用のダイヤ情報(当日から前後一日、計三日間)を取得する。
func TestGetNextTrainTimeInfoText(t *testing.T) {

	TestMode = true
	TodayForTest = time.Date(2016, 04, 24, 23, 59, 58, 0, time.Local)

	actual := GetNextTrainTimeInfoText(3)

	count := len(actual)
	if count < 10 {
		t.Log("取得できた文字列")
		t.Log(actual)
		t.Errorf("GetNextTrainTimeInfoText() is faild.")
	}
}

