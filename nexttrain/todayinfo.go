package nexttrain

import (
	"net/http"
	"io/ioutil"
)


// CalenderAPIのURL
// const date_api_url string = "http://calendar-service.net/cal?start_year=2010&start_mon=1&end_year=&end_mon=&year_style=normal&month_style=ja&wday_style=ja&format=json"
const date_api_url string = "http://calendar-service.net/cal?start_year=2010&start_mon=1&end_year=&end_mon=&year_style=normal&month_style=numeric&wday_style=en&format=csv"

func GetCalenderInfoJson() string {

	resp, _ := http.Get(date_api_url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)


	return string(byteArray) // htmlをstringで取得
}
