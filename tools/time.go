package tools

import "time"

type Time struct {
}

// DaysBetween 计算两个时间之间的差值（天数）
func (*Time) DaysBetween(t1, t2 time.Time) int {
	return int(t2.Sub(t1).Hours() / 24)
}

// AddDays 给时间加上指定天数
func (*Time) AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}
