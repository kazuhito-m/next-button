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

// 日付によってURLを組み立てる。(土曜日版)
func TestMakeTrainTimeUrlSat(t *testing.T) {

	targetDate := time.Date(2001, 12, 29, 0, 0, 0, 0, time.Local)
	param := DateInfo{}
	param.Date = targetDate
	param.DayNoOfWeek = 6
	param.SpecialDay = false

	actual := MakeTrainTimeUrl(param)

	// 土曜日のはずなので、URLのサフィックスは土曜用のはず
	r := regexp.MustCompile(".*_sat.htm")
	if !r.MatchString(actual) {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeTrainTimeUrl() is faild.URL syntax error.")
	}
}

// 日付によってURLを組み立てる。(日曜日版)
func TestMakeTrainTimeUrlSun(t *testing.T) {

	targetDate := time.Date(2001, 12, 30, 0, 0, 0, 0, time.Local)
	param := DateInfo{}
	param.Date = targetDate
	param.DayNoOfWeek = 0
	param.SpecialDay = false

	actual := MakeTrainTimeUrl(param)

	// 日曜日のはずなので、URLのサフィックスは休日用のはず
	r := regexp.MustCompile(".*_holi.htm")
	if !r.MatchString(actual) {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeTrainTimeUrl() is faild.URL syntax error.")
	}
}

// 日付によってURLを組み立てる。(平日だが祝日版)
func TestMakeTrainTimeUrlHoli(t *testing.T) {

	targetDate := time.Date(2002, 1, 1, 0, 0, 0, 0, time.Local)
	param := DateInfo{}
	param.Date = targetDate
	param.DayNoOfWeek = 2
	param.SpecialDay = true

	actual := MakeTrainTimeUrl(param)

	// 月曜日だが祝日のはずなので、URLのサフィックスは休日用のはず
	r := regexp.MustCompile(".*_holi.htm")
	if !r.MatchString(actual) {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeTrainTimeUrl() is faild.URL syntax error.")
	}
}
