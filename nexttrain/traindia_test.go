package nexttrain

import (
	"testing"
	"time"
	"strings"
)



// 日付によってURLを組み立てる。
func TestMakeTrainTimeUrl(t *testing.T) {

	targetDate := time.Date(2001, 12, 31, 0, 0, 0, 0, time.Local)
	param := DateInfo{}
	param.Date = targetDate

	actual := MakeTrainTimeUrl(targetDate)
	
	if if len(actual) <= 0 {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeTrainTimeUrl() is faild.")
	}
}

