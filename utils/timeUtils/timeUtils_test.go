package timeUtils

import (
	"fmt"
	"testing"
	"time"
)

/**
 * 时间戳的转换
 */
func TestUnixTime(t *testing.T) {
	time := TimeConfig{time.Now()}
	unixMilli := time.UnixMilli()
	fmt.Println("time to unix: ", unixMilli)
	fmt.Println("unixMilli to time, Year=", time.UGetYear(unixMilli))
}

/**
 * 时间格式的转换
 */
func TestTimeFormat(t *testing.T) {
	time := TimeConfig{time.Now()}
	fmt.Println("time=", time)
	fmt.Println("timeFormat DayMill=", time.FormatDayMill())
	fmt.Println("timeFormat Day=", time.FormatDay())
	fmt.Println("timeFormat Mill=", time.FormatMill())
	fmt.Println("timeFormat The custom=", time.Format("2006-01-02-15-04-05"))
}
