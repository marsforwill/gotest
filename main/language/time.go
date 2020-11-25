package main

import (
	"fmt"
	"time"
)

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetFirstDateOfLastMonth 获取传入的时间上个月份的第一天
func GetFirstDateOfLastMonth(d time.Time) time.Time {
	d = d.AddDate(0, -1, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的23:59:59
func GetLastDateOfLastMonth(d time.Time) time.Time {
	return GetEndTime(GetFirstDateOfMonth(d).AddDate(0, 0, -1))
}

// GetHalfDateOfMonth 15号0点
func GetHalfDateOfMonth(d time.Time) time.Time {
	return GetZeroTime(GetFirstDateOfMonth(d).AddDate(0, 0, 14))
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetEndTime 获取某一天的23:59:59
func GetEndTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

func main() {
	fmt.Println(GetFirstDateOfLastMonth(time.Now()))
	fmt.Println(GetLastDateOfLastMonth(time.Now()))
	fmt.Println(GetFirstDateOfMonth(time.Now()))
	//fmt.Println(time.Unix(1606752000,0).AddDate(0, -1, 0))
	fmt.Println(GetFirstDateOfMonth(time.Now()))
	fmt.Println(GetHalfDateOfMonth(time.Now()))
}
