package schedule

import "tlan/plan"

type Track struct {
	Schedule Schedule
	Name     string
	Projects []*plan.Project
}

type Schedule struct {
	Name   string
	Period Period
}

type Time struct {
	Hour   int
	Minute int
}

type Period struct {
	Start Time
	End   Time
}

func FilterTracks(arr []*Track, cond func(track Track) bool) []*Track {
	var result []*Track
	for i := range arr {
		if cond(*arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
