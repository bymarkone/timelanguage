package schedule

import (
	"tlan/plan"
	"tlan/utils"
)

type Track struct {
	Schedule Schedule
	Name     string
	Projects []*plan.Project
}

type Schedule struct {
	Name   string
	Period utils.Period
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
