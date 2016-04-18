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

// 鉄道のダイヤ情報の一つを表す構造体
type TrainTimeInfo struct {
	TrainTime       time.Time
	Destination 	int
}


// 「次の電車」情報の束を取得する。
func GetNextTrainTimeInfo() TrainTimeInfo[] {

}

// 「指定日の時刻情報」を束で取得する。
func GetTrainTimeInfo(targetDate DateInfo) TrainTimeInfo[]() {

}

// URLで取得したページの文字列(HTML)から、時刻表情報を取得する。
func ScrapeTrainTimeInfo(url string) TrainTimeInfo[]() {
	return new TrainTimeInfo[]
}

// 指定された日付情報から、時刻表のページのURLを組み立てる。
func MakeTrainTimeUrl(targetDate DateInfo) string {
	return "仮実装"
}
