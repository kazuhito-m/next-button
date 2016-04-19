package nexttrain

import (
	"testing"
	"time"
	"regexp"
)



// 日付によってURLを組み立てる。
func TestMakeTrainTimeUrl(t *testing.T) {

	targetDate := time.Date(2001, 12, 31, 0, 0, 0, 0, time.Local)
	param := DateInfo{}
	param.Date = targetDate
	param.DayNoOfWeek = 1
	param.SpecialDay = false

	actual := MakeTrainTimeUrl(param)

	if len(actual) <= 0 {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeTrainTimeUrl() is faild.")
	}

	// 平日のはずなので、URLのサフィックスは"数字数桁.htm"のはず
	r := regexp.MustCompile(".*\\d{4}\\.htm")
	if !r.MatchString(actual) {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeTrainTimeUrl() is faild.URL syntax error.")
	}
}

