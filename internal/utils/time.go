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

func (p DateTime) ToString() string {
	return strconv.Itoa(p.Hour)
}

func (dt DateTime) toTime() time.Time {
	return time.Date(dt.Year, dt.Month, dt.Day, dt.Hour, dt.Minute, 0, 0, time.Now().Location())
}

type Period struct {
	Weekdays []time.Weekday
	Start    DateTime
	End      DateTime
}

func (p Period) ToString() string {
	results := ""
	if len(p.Weekdays) == 7 {
		results = "daily"
	} else {
		for _, weekday := range p.Weekdays {
			results = results + weekday.String()[0:3]
		}
	}
	return results + ", " + p.Start.ToString() + "-" + p.End.ToString()
}

func (p Period) ActiveIn(date time.Time) bool {
	return p.Start.toTime().Before(date) && p.End.toTime().After(date) || p.Start.toTime().Equal(date)
}

func Parse(value string) (int, int, int) {
	first, _ := strconv.Atoi(value[0:2])
	second, _ := strconv.Atoi(value[3:5])
	third := 0
	if len(value) > 5 {
		third, _ = strconv.Atoi("20" + value[6:8])
	}
	return first, second, third
}
