package goutil

import (
	"fmt"
	"strconv"
	"time"
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func TimeStampParseString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
func NowInt() int64 {
	return time.Now().Unix()
}

func TimeStringParseTimeStamp(timeString string) int64 {
	return TimeStringParseTime(timeString).Unix()
}

func TimeStringParseTime(timeString string) time.Time {
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeString, local)
	return t
}

func TimeParseString(time2 time.Time) string {
	return time2.Format("2006-01-02 15:04:05")
}
func NextMonths(timeString string, months int) string {
	nowTime := TimeStringParseTime(timeString)
	getTime := nowTime.AddDate(0, months, 0)
	return TimeParseString(getTime)
}
func NextDays(timeString string, days int) string {
	nowTime := TimeStringParseTime(timeString)
	getTime := nowTime.AddDate(0, 0, days)
	return TimeParseString(getTime)
}
func NextYears(timeString string, years int) string {
	nowTime := TimeStringParseTime(timeString)
	getTime := nowTime.AddDate(years, 0, 0)
	return TimeParseString(getTime)
}
func NowFloat() float64 {
	value := int(time.Now().UnixNano()) / 1e6
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(value)/1000), 64)
	res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	return res
}
