package nexttrain

import "testing"

// カレンダーAPIから内容が取れるか
// TODO 外部に接続するテストなんで、外部依存をなくすよう改造したい。
func TestGetCalenderInfoJson(t *testing.T) {
	actual := GetCalenderInfoJson()
	if len(actual) <= 0 || true {
		t.Log("取得できた文字列")
		t.Log(actual)
		t.Errorf("GetCalenderInfoJson() is faild.")
	}
}

