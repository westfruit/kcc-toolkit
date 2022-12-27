package convert

import (
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

// 时间戳转换为日期字符串
func TimestampToStr(sec int64) string {
	if sec == 0 {
		return ""
	}

	tm := time.Unix(sec, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func NowToDateTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 时间戳转换为日期字符串
func TimestampToStrDate(sec int64) string {
	if sec == 0 {
		return ""
	}

	tm := time.Unix(sec, 0)
	return tm.Format("2006-01-02")
}

// 时间字符串转化为时间戳
func StrToTimestamp(str string) int64 {
	// 如果不包含时间，则加上时间段
	if len(str) == 10 {
		str = str + " 00:00:00"
	}

	//loc, _ := time.LoadLocation("Local") //获取时区
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		logrus.Error("StrToTimestamp转换错误, str=", str, ", ", err)
		return 0
	}

	return t.Unix()
}

// 14位格式（yyyyMMddHHmmss）字符串转换为时间戳
func YyyyMMddHHmmssToTimestamp(str string) int64 {

	if len(str) != 14 {
		return 0
	}

	t, err := time.ParseInLocation("20060102150405", str, time.Local)
	if err != nil {
		logrus.Error("YyyyMMddHHmmssToTimestamp转换错误, str=", str, ", err=", err)
		return 0
	}

	return t.Unix()
}

// 时间点转换为时间戳
func TimePointToTimestamp(timePoint string) int64 {

	layout := "2006-01-02 15:04" //转化所需模板
	date := time.Now().Format("2006-01-02")
	value := date + " " + timePoint                        //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的
	loc, _ := time.LoadLocation("Local")                   //获取时区
	theTime, _ := time.ParseInLocation(layout, value, loc) //使用模板在对应时区转化为time.time类型

	return theTime.Unix()
}

func GetCurrentDateStr() string {
	return time.Now().Format("2006-01-02")
}

// Format("20060102")
func GetCurrentDateNumber() string {
	return time.Now().Format("20060102")
}

func GetCurrentDateTimeNumber() string {
	return time.Now().Format("20060102150405")
}

//获得今天0时0分0秒时间戳
func GetTodayStartTimestamp() int64 {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return tm1.Unix()
}

//获得今天23时59分59秒时间戳
func GetTodayEndTimestamp() int64 {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999, t.Location())

	return tm1.Unix()
}

// 获得明天0时0分0秒时间戳
func GetTomorrowStartTimestamp() int64 {
	t := time.Now().AddDate(0, 0, 1)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return tm1.Unix()
}

// 获得昨天0时0分0秒时间戳
func GetYesterdayStartTimestamp() int64 {
	t := time.Now().AddDate(0, 0, -1)
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return tm1.Unix()
}

// 获取2个时间戳之间的天数
func GetDaysBetweenTwoTimestamp(start int64, end int64) float64 {
	c := math.Abs(float64(end - start))

	days := c / (24 * 60 * 60)

	return days
}

func DateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func DateNumberString(t time.Time) string {
	return t.Format("20060102")
}

func DateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func RFC3339ToInt64(value string) int64 {

	a, err := RFC3339ToCSTLayout(value)
	if err != nil {
		return 0
	}
	return StrToTimestamp(a)
}

//将 2020-11-08T08:18:46+08:00 转成 2020-11-08 08:18:46
func RFC3339ToCSTLayout(value string) (string, error) {

	const CSTLayout = "2006-01-02 15:04:05"
	var cst *time.Location
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

// 获取指定时间戳的0时0分0秒时间戳
func GetStartTimestamp(sec int64) int64 {
	if sec <= 0 {
		return 0
	}

	t := time.Unix(sec, 0)
	tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	return tm.Unix()
}

// 获取指定时间戳的23时59分59秒时间戳
func GetEndTimestamp(sec int64) int64 {
	if sec <= 0 {
		return 0
	}

	t := time.Unix(sec, 0)
	tm := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999, t.Location())

	return tm.Unix()
}

func Int64Totime(value int64) time.Time {
	const CSTLayout = "2006-01-02 15:04:05"
	var cst *time.Location
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	a := TimestampToStr(value)
	ts, err := time.Parse(CSTLayout, a)
	if err != nil {
		return time.Now()
	}

	return ts.In(cst)
}

//  检查int64是秒还是毫秒, 返回秒
func CheckInt64ForTimestamp(value int64) int64 {
	if value < 0 {
		return 0
	}

	// 检查int64是秒还是毫秒
	if value > 999999999 {
		return value / 1000
	}

	return value
}

// 获取当前时间戳
func GetNowTimestamp() int64 {
	return time.Now().Unix()
}
