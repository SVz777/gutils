/**
 * @file    time.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020-02-06
 * @desc
 */
package time

import (
	"errors"
	"github.com/SVz777/gutils/convert"
	"time"
)

const (
	DefaultTime   = "1971-01-01 00:00:00"
	LayoutTime    = "2006-01-02 15:04:05"
	LayoutTimeYMD = "2006-01-02"
	LayoutTimeYM  = "2006-01"
	LayoutTimeYmd = "20060102"
	LayoutTimeYm  = "200601"
	LayoutTimeY   = "2006"
	LayoutTimeHI  = "15:04"

	ChineseLayoutTime = "2006年1月2日 15:04:05"
	ChineseLayoutMD   = "01月02日"
)

//获取当前时间 YYYY-MM-DD H:i:s
func GetCurrentDateTime() (currentTime string) {
	return time.Now().Format(LayoutTime)
}

//获取当前中文时间 YYYY年MM月DD日 H:i:s
func GetCurrentChineseDateTime() (currentTime string) {
	return time.Now().Format(ChineseLayoutTime)
}
func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(LayoutTime)
}

//时间戳转换为date，输出指定格式
func TimestampToDateWithLayout(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

//获取指定的年月日 YYYY年MM月DD日 ，+/- time
func GetDayDateTimeWithExtra(extra string) string {
	curTime := time.Now() //获取系统当前时间
	dh, _ := time.ParseDuration(extra)
	return curTime.Add(dh).Format(LayoutTimeYMD)
}

//获取指定的年月日 YYYY年MM月DD日 ，+/- time
func GetDayDateTimeWithExtraLayout(extra string, layout string, date string) (string, error) {
	afterDate, err := time.Parse(LayoutTime, date)
	if err != nil {
		return "", err
	}
	dh, _ := time.ParseDuration(extra)
	return afterDate.Add(dh).Format(layout), nil
}

func GetDayStart(date string, extra string) (result string, err error) {
	datetime, err := time.Parse(LayoutTime, date)
	if nil != err {
		return "", err
	}
	targetDatetime := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 0, 0, 0, 0, datetime.Location())
	dh, _ := time.ParseDuration(extra)
	return targetDatetime.Add(dh).Format(LayoutTime), nil
}

func GetDayEnd(date string, extra string) (result string, err error) {
	datetime, err := time.Parse(LayoutTime, date)
	if nil != err {
		return "", err
	}
	targetDatetime := time.Date(datetime.Year(), datetime.Month(), datetime.Day(), 23, 59, 59, 0, datetime.Location())
	dh, _ := time.ParseDuration(extra)
	return targetDatetime.Add(dh).Format(LayoutTime), nil
}

func DatetimeToTimestamp(date string, layout string) (timestamp int64, err error) {
	if layout == "" {
		layout = LayoutTime
	}
	datetime, err := time.Parse(layout, date)
	if nil != err {
		return 0, err
	}
	return datetime.Unix(), nil
}

func TimestampToDatetime(timestamp int64, layout string) (datetime string, err error) {
	if layout == "" {
		layout = LayoutTime
	}
	time := time.Unix(timestamp, 0)
	return time.Format(layout), nil
}

func ParseTime(dateTime string) (time.Time, error) {
	if t, err := time.ParseInLocation(LayoutTime, dateTime, time.Local); err == nil {
		return t, nil
	}
	if t, err := time.ParseInLocation(LayoutTimeYMD, dateTime, time.Local); err == nil {
		return t, nil
	}
	if t, err := time.ParseInLocation(LayoutTimeYM, dateTime, time.Local); err == nil {
		return t, nil
	}
	if t, err := time.ParseInLocation(LayoutTimeY, dateTime, time.Local); err == nil {
		return t, nil
	}
	return time.Now(), errors.New("parsing time error:" + dateTime)
}

func GetMonthDays(year, month int) int {
	d1 := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	d2 := d1.AddDate(0, 1, -1)
	return d2.Day() - d1.Day() + 1
}

func GetYearMonth(yearMonthS string) int32 {
	d, _ := ParseTime(yearMonthS)
	ymS := d.Format(LayoutTimeYm)
	ym, _ := convert.Int32(ymS)
	return ym
}

func GetYearMonthDay(yearMonthdayS string) int32 {
	d, _ := ParseTime(yearMonthdayS)
	ymdS := d.Format(LayoutTimeYmd)
	ymd, _ := convert.Int32(ymdS)
	return ymd
}

func GetYearMonthString(yearMonth int32) string {
	ym, _ := convert.String(yearMonth)
	d, _ := time.ParseInLocation(LayoutTimeYm, ym, time.Local)
	return d.Format(LayoutTimeYM)
}
