package timeUtils

import (
	"strconv"
	"time"
)

type TimeConfig struct {
	Time time.Time
}

var (
	dayMill = "2006-01-02 15:04:05"
	day     = "2006-01-02"
	mill    = "15:04:05"
)

/**
 * 输出年月日
 */
func (t *TimeConfig) Year() int {
	return t.Time.Year()
}
func (t *TimeConfig) Month() int {
	return int(t.Time.Month())
}
func (t *TimeConfig) Day() int {
	return t.Time.Day()
}
func (t *TimeConfig) Hour() int {
	return t.Time.Hour()
}
func (t *TimeConfig) Minute() int {
	return t.Time.Minute()
}
func (t *TimeConfig) Second() int {
	return t.Time.Second()
}

/**
 * 输出格式化时间
 * 例如："2006-01-02"
 */
func (t *TimeConfig) FormatDayMill() string {
	format := t.Time.Format(dayMill)
	return format
}

/**
 * 输出格式化时间
 * 例如："15:04:05"
 */
func (t *TimeConfig) FormatMill() string {
	format := t.Time.Format(mill)
	return format
}

/**
 * 输出格式化时间
 * 例如："2006-01-02 15:04:05"
 */
func (t *TimeConfig) FormatDay() string {
	format := t.Time.Format(day)
	return format
}

/**
 * 格式化时间
 * 例如："2006-01-02-15-04-05"
 */
func (t *TimeConfig) Format(layout string) string {
	format := t.Time.Format(layout)
	return format
}

// 获取以秒作为单位的时间戳
func (t *TimeConfig) Unix() int64 {
	unix := t.Time.Unix()
	return unix
}

// 获取以毫秒作为单位的时间戳
func (t *TimeConfig) UnixMilli() int64 {
	unix := t.Time.UnixMilli()
	return unix
}

//获取时间戳所在年
func (t *TimeConfig) UGetYear(timeStamp int64) int {
	thisTime := time.Unix(timeStamp, 0).In(time.FixedZone("CST", 8*3600))
	if i := len(strconv.FormatInt(timeStamp, 10)); i == 13 {
		thisTime = time.UnixMilli(timeStamp)
	}
	return thisTime.Year()
}

//获取时间戳所在月
func (t *TimeConfig) UGetMonth(timeStamp int64) int {
	thisTime := time.Unix(timeStamp, 0).In(time.FixedZone("CST", 8*3600))
	if i := len(strconv.FormatInt(timeStamp, 10)); i == 13 {
		thisTime = time.UnixMilli(timeStamp)
	}
	return int(thisTime.Month())
}

//获取时间戳所在日
func (t *TimeConfig) UGetDay(timeStamp int64) int {
	thisTime := time.Unix(timeStamp, 0).In(time.FixedZone("CST", 8*3600))
	if i := len(strconv.FormatInt(timeStamp, 10)); i == 13 {
		thisTime = time.UnixMilli(timeStamp)
	}
	return thisTime.Day()
}
