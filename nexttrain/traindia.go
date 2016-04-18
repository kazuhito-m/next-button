package nexttrain

import (
	"time"
)

// 鉄道のダイヤ情報の一つを表す構造体
type TrainTimeInfo struct {
	TrainTime       time.Time
	Destination 	int
}


// 「次の電車」情報の束を取得する。
func GetNextTrainTimeInfo() []TrainTimeInfo {
	return []TrainTimeInfo{}
}

// 「指定日の時刻情報」を束で取得する。
func GetTrainTimeInfo(targetDate DateInfo) []TrainTimeInfo {
	// 仮実装
	println(targetDate.Date.String())
	return []TrainTimeInfo{}
}

// URLで取得したページの文字列(HTML)から、時刻表情報を取得する。
func ScrapeTrainTimeInfo(url string) []TrainTimeInfo {
	println(url)
	return []TrainTimeInfo{}
}

// 指定された日付情報から、時刻表のページのURLを組み立てる。
func MakeTrainTimeUrl(targetDate DateInfo) string {
	return "仮実装"
}
