package nexttrain

import (
	"testing"
	"time"
	"strings"
)


// テスト用Csvデータ
// TODO ファイル読み込みにしたい…
const Testcsv string = `年,月,日,年号,和暦,曜日,曜日番号,祝日名
2010,01,01,平成,22,Fri,5,元日
2010,01,02,平成,22,Sat,6,
2010,01,03,平成,22,Sun,0,
2010,01,04,平成,22,Mon,1,
2010,01,05,平成,22,Tue,2,
2010,01,06,平成,22,Wed,3,
2010,01,07,平成,22,Thu,4,
2010,01,08,平成,22,Fri,5,
2010,01,09,平成,22,Sat,6,
2010,01,10,平成,22,Sun,0,
2010,01,11,平成,22,Mon,1,成人の日
2010,01,12,平成,22,Tue,2,
2010,01,13,平成,22,Wed,3,
2010,01,14,平成,22,Thu,4,
2010,01,15,平成,22,Fri,5,
2010,01,16,平成,22,Sat,6,
2010,01,17,平成,22,Sun,0,
2010,01,18,平成,22,Mon,1,
2010,01,19,平成,22,Tue,2,
2010,01,20,平成,22,Wed,3,
2010,01,21,平成,22,Thu,4,
2010,01,22,平成,22,Fri,5,
2010,01,23,平成,22,Sat,6,
2010,01,24,平成,22,Sun,0,
2010,01,25,平成,22,Mon,1,
2010,01,26,平成,22,Tue,2,
2010,01,27,平成,22,Wed,3,
2010,01,28,平成,22,Thu,4,
2010,01,29,平成,22,Fri,5,
2010,01,30,平成,22,Sat,6,
2010,01,31,平成,22,Sun,0,
`

// 日付によってURLを組み立てる。
func TestMakeApiUrlByDate(t *testing.T) {
	targetDate := time.Date(2001, 12, 31, 0, 0, 0, 0, time.Local)
	actual := MakeApiUrlByDate(targetDate)
	if !strings.Contains(actual, "start_year=2001") || !strings.Contains(actual, "start_mon=12") {
		t.Log("作成したURL")
		t.Log(actual)
		t.Errorf("MakeApiUrlByDate() is faild.")
	}
}

// カレンダーAPIから内容が取れるか
// TODO 外部に接続するテストなんで、外部依存をなくすよう改造したい。
func TestGetCalenderInfoJson(t *testing.T) {
	actual := GetCalenderInfoCsv(time.Now())
	if len(actual) <= 0 {
		t.Log("取得できた文字列")
		t.Log(actual)
		t.Errorf("GetCalenderInfoJson() is faild.")
	}
}

// カレンダーAPいから取り出したCSVから構造体の配列へ変換する
func TestConvDateInfosByCsv(t *testing.T) {

	actual := ConvDateInfosByCsv(Testcsv)

	// 件数確認
	if len(actual) != 31 {
		t.Errorf("ConvDateInfosByCsv() is faild. record count fail.")
	}

	// 一件目の特徴
	di := actual[0]
	if !di.Date.Equal(time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local)) {
		t.Errorf("ConvDateInfosByCsv() is faild. first record day not equals.")
	}
	if di.DayNoOfWeek != 5 {
		t.Errorf("ConvDateInfosByCsv() is faild. first record weekday not equals.")
	}
	if !di.SpecialDay {
		t.Errorf("ConvDateInfosByCsv() is faild. first record SpecialDay not equals.")
	}

	// 最終件の特徴
	di = actual[30]
	if !di.Date.Equal(time.Date(2010, 1, 31, 0, 0, 0, 0, time.Local)) {
		t.Errorf("ConvDateInfosByCsv() is faild. last record day not equals.")
	}
	if di.DayNoOfWeek != 0 {
		t.Errorf("ConvDateInfosByCsv() is faild. last record weekday not equals.")
	}
	if di.SpecialDay {
		t.Errorf("ConvDateInfosByCsv() is faild. last record SpecialDay not equals.")
	}

}

// 日付だけ比較。同じになる場合
func TestEqualDateOnly(t *testing.T) {

	t1 := time.Date(2014, 12, 20, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2014, 12, 20, 23, 59, 59, 0, time.Local)

	actual := EqualDateOnly(t1, t2)

	if !actual {
		t.Errorf("EqualDateOnly() is faild.")
	}
}

// 日付だけ比較。異なる場合
func TestEqualDateOnlyNotEqual(t *testing.T) {

	t1 := time.Date(2014, 12, 20, 23, 59, 58, 0, time.Local)
	t2 := time.Date(2014, 12, 21, 0, 0, 0, 0, time.Local)

	actual := EqualDateOnly(t1, t2)

	if actual {
		t.Errorf("EqualDateOnly() no2 is faild.")
	}
}

func TestGetDay(t *testing.T) {

	TestMode = true
	TodayForTest = time.Date(2014, 12, 20, 23, 59, 58, 0, time.Local)

	actual := GetDay(-1)

	if actual.Year() == 2014 && actual.Month() == 12 && actual.Day() == 19 {
		t.Errorf("GetDay() is faild. -1.")
	}

	actual = GetDay(1)

	if actual.Year() == 2014 && actual.Month() == 12 && actual.Day() == 21 {
		t.Errorf("GetDay() is faild. +1.")
	}

}

func TestGetTodayInfo(t *testing.T) {

	TestMode = true
	TodayForTest = time.Date(2016, 04, 24, 23, 59, 58, 0, time.Local)

	actual := GetTodayInfo()

	if actual.Date.Day() != 24 {
		t.Errorf("TestGetTodayInfo() is faild. Day is not 24")
	}

	if actual.DayNoOfWeek != 0 {
		t.Errorf("TestGetTodayInfo() is faild. Week day number not Sunday.")
	}

	if actual.SpecialDay {
		t.Errorf("TestGetTodayInfo() is faild. Special day is faild.")
	}

}


// 日の情報を取得する(祝日版)
func TestGetTodayInfoForHolyday(t *testing.T) {

	TestMode = true
	TodayForTest = time.Date(2016, 04, 24, 23, 59, 58, 0, time.Local)

	actual := GetDayInfo(5)    // 昭和の日設定

	if actual.Date.Day() != 29 {
		t.Errorf("TestGetTodayInfo() is faild. Day is not 29")
	}

	if actual.DayNoOfWeek != 5 {
		t.Errorf("TestGetTodayInfo() is faild. Week day number not Sunday.")
	}

	if !actual.SpecialDay {
		t.Errorf("TestGetTodayInfo() is faild. Special day is faild.")
	}

}