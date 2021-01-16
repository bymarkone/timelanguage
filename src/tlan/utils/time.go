package utils

import (
	"strconv"
	"time"
)

type DateTime struct {
	Minute int
	Hour   int
	Day    int
	Month  time.Month
	Year   int
}

func (dt DateTime) toTime() time.Time {
	return time.Date(2021, dt.Month, dt.Day, dt.Hour, dt.Minute, 0, 0, time.Now().Location())
}

type Period struct {
	Start DateTime
	End   DateTime
}

func (p Period) ActiveIn(date time.Time) bool {
	return p.Start.toTime().Before(date) && p.End.toTime().After(date) || p.Start.toTime().Equal(date)
}

func Parse(value string) (int, int) {
	first, _ := strconv.Atoi(value[0:2])
	second, _ := strconv.Atoi(value[2:4])
	return first, second
}
