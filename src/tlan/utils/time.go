package utils

import "strconv"

type DateTime struct {
	Minute int
	Hour   int
	Day    int
	Month  int
	Year   int
}

type Period struct {
	Start DateTime
	End   DateTime
}

func Parse(value string) (int, int) {
	first, _ := strconv.Atoi(value[0:2])
	second, _ := strconv.Atoi(value[2:4])
	return first, second
}

