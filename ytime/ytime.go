package ytime

import "time"

const (
	hourSec = 3600
	fmtStr  = "2006-01-02 15:04:05"
)

// UnixToStr 秒级时间戳转时间字符串
func UnixToStr(unix int64) string {
	return time.Unix(unix, 0).Format(fmtStr)
}
