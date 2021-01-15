package utils

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