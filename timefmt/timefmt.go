package timefmt

import "time"

//GetUnixTimeByFmt 根据指定时间格式获取对应时间戳
/**
eg: "2006-01-02": 获取当天零点
"2006-01": 获取当月第一天零点
@return 时间戳
*/
func GetUnixTimeByFmt(fmtstr string) int64 {
	curTimestr := time.Now().Format(fmtstr)
	t, _ := time.ParseInLocation(fmtstr, curTimestr, time.Local)

	return t.Unix()
}
