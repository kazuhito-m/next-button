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
	if len(actual) != 31 {
		t.Errorf("ConvDateInfosByCsv() is faild.")
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

