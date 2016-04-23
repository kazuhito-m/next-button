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

// とある日の平日ダイヤを取得する。
func TestScrapeTrainTimeInfoBasic(t *testing.T) {
	// とある平日のダイヤを返すURL
	const url = "http://www.ekikara.jp/newdata/ekijikoku/2701062/down1_27212011.htm"
	targetDate := time.Date(2002, 1, 1, 0, 0, 0, 0, time.Local)

	actual := ScrapeTrainTimeInfo(url, targetDate)

	count := len(actual)
	if count != 92 {
		t.Log("取得できたダイヤの要素数")
		t.Log(count)
		t.Errorf("ScrapeTrainTimeInfo() is faild.dia count %d but %d", 92, count)
	}
}

// とある日の平日ダイヤを日付していだけで取得する。
func TestGetTrainTimeInfo(t *testing.T) {

	targetDate := time.Date(2001, 12, 31, 0, 0, 0, 0, time.Local)
	param := DateInfo{}
	param.Date = targetDate
	param.DayNoOfWeek = 1
	param.SpecialDay = false

	actual := GetTrainTimeInfo(param)

	count := len(actual)
	if count != 92 {
		t.Log("取得できたダイヤの要素数")
		t.Log(count)
		t.Errorf("GetTrainTimeInfo() is faild.dia count %d but %d", 92, count)
	}
}

// 判定用のダイヤ情報(当日から前後一日、計三日間)を取得する。
func TestGetTrainTimeInfoFullRange(t *testing.T) {

	TestMode = true
	TodayForTest = time.Date(2016, 04, 24, 23, 59, 58, 0, time.Local)

	actual := GetTrainTimeInfoFullRange()

	const expect = 266
	count := len(actual)
	if count != expect {
		t.Log("取得できたダイヤの要素数")
		t.Log(count)
		t.Errorf("GetTrainTimeInfo() is faild.dia count %d but %d", expect, count)
	}
}

// 「現在」より後の「指定した個数分」のダイヤを取ってくる
func TestGetNextTrainTimeInfo(t *testing.T) {

	const expect = 5

	TestMode = true
	TodayForTest = time.Date(2016, 04, 23, 23, 30, 58, 0, time.Local)

	actual := GetNextTrainTimeInfo(expect)

	count := len(actual)
	if count != expect {
		t.Log("取得できたダイヤの要素数")
		t.Log(count)
		t.Errorf("GetNextTrainTimeInfo() is faild.dia count %d but %d", expect, count)
	}
}
