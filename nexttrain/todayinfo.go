package nexttrain

import (
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
	"encoding/csv"
	"strings"
	"io"
	"strconv"
	"unicode/utf8"
)


// CalenderAPIのURL
const date_api_url string = "http://calendar-service.net/cal?start_year=%d&start_mon=%d&end_year=&end_mon=&year_style=normal&month_style=numeric&wday_style=en&format=csv"

// 日付情報を入れる構造体
type DateInfo struct {
	Date        time.Time
	DayNoOfWeek int
	SpecialDay  bool
}


func GetCalenderInfoCsv(target time.Time) string {

	// 対象日(というより対象日の月)を指定したURL作成
	url := MakeApiUrlByDate(target)
	// HTTP通信にて、CSVを取得
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	return string(byteArray) // htmlをstringで取得
}

// 指定した日付のカレンダーAPIのURLを作成し取得
func MakeApiUrlByDate(target time.Time) string {
	return fmt.Sprintf(date_api_url, target.Year(), target.Month())
}

func ConvDateInfosByCsv(csvText string) []DateInfo {
	// 配列を作成
	dateInfos := []DateInfo{}
	// 一行ずつカンマ区切りを処理
	csvReader := csv.NewReader(strings.NewReader(csvText))

	isFirst := true
	for {
		record, err := csvReader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if isFirst {
			isFirst = false
			continue
		}

		// 無事読めたようだ。日付作成。
		year, _ := strconv.Atoi(record[0])
		mon, _ := strconv.Atoi(record[1])
		day, _ := strconv.Atoi(record[2])
		dayNoOfWeek, _ := strconv.Atoi(record[6])
		date := time.Date(year, time.Month(mon), day, 0, 0, 0, 0, time.Local)
		// 構造体作成
		ti := DateInfo{}
		ti.Date = date
		ti.DayNoOfWeek = dayNoOfWeek   // 曜日
		ti.SpecialDay = (utf8.RuneCountInString(record[7]) > 0)
		// 配列に追加
		dateInfos = append(dateInfos, ti)
	}
	return dateInfos
}


// 日付から時刻を取り去る
func TruncateHms(src time.Time) time.Time {
	return src.Truncate(time.Hour).Add(- time.Duration(src.Hour()) * time.Hour)

}

// ２つの時刻変数の日付だけを取り出して同一か判定する
func EqualDateOnly(src time.Time, dest time.Time) bool {
	return TruncateHms(src).Equal(TruncateHms(dest))
}
