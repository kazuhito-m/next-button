package nexttrain
import "fmt"

// 「言語化」に関するモジュール。
// 万が一まがりまちがって「他言語化」とかされるときは、ここを弄う。

func GetNextTrainTimeInfoText(count int) string {
	infos := GetNextTrainTimeInfo(count)
	// 一個一個の列車情報を作る
	var infoStrs string = ""
	for _ , info := range infos {
		infoStrs = infoStrs + makeTrainInfoStr(info)
	}
	return fmt.Sprintf("次の列車は、%sです。" , infoStrs)
}

func makeTrainInfoStr(info TrainTimeInfo)string {
	tTime := info.TrainTime
	return fmt.Sprintf("約%d分後、%d時%d分、", info.Destination , tTime.Hour(),tTime.Minute())
}
