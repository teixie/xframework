package xtime

import (
	"fmt"
	"regexp"
	"time"
)

const (
	FmtHi     = "15:04"
	Fmtmd     = "01-02"
	FmtHis    = "15:04:05"
	FmtYmd    = "2006-01-02"
	FmtYmdHi  = "2006-01-02 15:04"
	FmtYmdHis = "2006-01-02 15:04:05"
)

var (
	location *time.Location
)

// 设置时区
func SetLocation(loc *time.Location) {
	location = loc
}

// 获取时区
func GetLocation() *time.Location {
	if location != nil {
		return location
	}
	return time.Local
}

// 获取当前时间
func Now() time.Time {
	return time.Now().In(GetLocation())
}

// 获取今天的起始时间
func Today() time.Time {
	t := Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetLocation())
}

// 获取今天的结束时间
func EndOfToday() time.Time {
	t := Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, GetLocation())
}

// 获取明天的起始时间
func Tomorrow() time.Time {
	t := Now().Add(24 * time.Hour)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetLocation())
}

// 获取昨天的起始时间
func Yesterday() time.Time {
	t := Now().Add(-24 * time.Hour)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetLocation())
}

// 获取某一天的起始时间
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetLocation())
}

// 获取某一天的结束时间
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, GetLocation())
}

// 获取Unix开始时间
func UnixStartTime() time.Time {
	return time.Date(1970, 1, 1, 0, 0, 1, 0, GetLocation())
}

// 解析时间
func ParseTime(t interface{}) (time.Time, error) {
	if t == nil {
		return time.Time{}, nil
	}

	switch t.(type) {
	case time.Time:
		return t.(time.Time), nil
	case string:
		return time.Parse(FmtYmdHis, t.(string))
	case int:
		return time.Unix(int64(t.(int)), 0), nil
	case int64:
		return time.Unix(t.(int64), 0), nil
	}

	return time.Time{}, fmt.Errorf("parse error")
}

// 时间格式化，格式如：YYYY-mm-dd HH:ii:ss
func FormatTime(t time.Time, fmtStr string) string {
	exists, err := regexp.Match("[YymdHis]+", []byte(fmtStr))
	if err == nil && !exists {
		return t.Format(fmtStr)
	}

	timeStr := t.String()
	o := map[string]string{
		"Y+": timeStr[0:4],
		"y+": timeStr[2:4],
		"m+": timeStr[5:7],
		"d+": timeStr[8:10],
		"H+": timeStr[11:13],
		"i+": timeStr[14:16],
		"s+": timeStr[17:19],
	}
	for k, v := range o {
		re, _ := regexp.Compile(k)
		fmtStr = re.ReplaceAllString(fmtStr, v)
	}
	return fmtStr
}
