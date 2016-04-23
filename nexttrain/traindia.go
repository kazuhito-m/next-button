package nexttrain

import (
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

const train_dia_url string = "http://www.ekikara.jp/newdata/ekijikoku/2701062/down1_27212011%s.htm"


// 鉄道のダイヤ情報の一つを表す構造体
type TrainTimeInfo struct {
	TrainTime   time.Time
	Destination int
}


// 「次の電車」情報の束を取得する。
func GetNextTrainTimeInfo() []TrainTimeInfo {
	return []TrainTimeInfo{}
}

// 「指定日の時刻情報」を束で取得する。
func GetTrainTimeInfo(targetDate DateInfo) []TrainTimeInfo {
	// 日付からURLを作成
	url := MakeTrainTimeUrl(targetDate)
	// ダイヤデータを取得、分解し束で返す
	return ScrapeTrainTimeInfo(url, targetDate.Date)
}

// URLで取得したページの文字列(HTML)から、時刻表情報を取得する。
func ScrapeTrainTimeInfo(url string, targetDate time.Time) []TrainTimeInfo {
	infos := []TrainTimeInfo{}
	td := targetDate
	// URLからWEBページ取ってくる
	doc, _ := goquery.NewDocument(url)
	doc.Find(".lowBg01 > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		// 時、取得
		var hour int = 0
		s.Find(".lowBg06 > .l").Each(func(_ int, s3 *goquery.Selection) {
			var hourStr = s3.Text()
			hour, _ = strconv.Atoi(hourStr)
		})
		if hour == 0 {
			return
		}
		// 分、取得。
		s.Find(".ll").Each(func(_ int, s4 *goquery.Selection) {
			var min, _ = strconv.Atoi(s4.Text())
			tti := TrainTimeInfo{}
			tti.TrainTime = time.Date(td.Year(), td.Month(), td.Day(), hour, min, 0, 0, time.Local)
			infos = append(infos, tti)
		})
	})
	return infos
}

// 指定された日付情報から、時刻表のページのURLを組み立てる。
func MakeTrainTimeUrl(targetDate DateInfo) string {
	diaSufix := ""    //
	// まず、日曜or休日なら
	if (targetDate.SpecialDay || targetDate.DayNoOfWeek == 0) {
		diaSufix = "_holi"
	} else if (targetDate.DayNoOfWeek == 6) {
		diaSufix = "_sat"
	} // それ以外はなにもつけなくてよし。
	return fmt.Sprintf(train_dia_url, diaSufix)
}
