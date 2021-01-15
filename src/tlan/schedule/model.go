package schedule

import (
	plan "tlan/plan"
	"tlan/utils"
)

type Track struct {
	Schedule      Schedule
	Name          string
	Projects      []*plan.Project
	_cachedActive []*plan.Project
}

func (t *Track) ActiveProjects() []*plan.Project {
	if len(t._cachedActive) == 0 {
		t._cachedActive = plan.FilterProjects(t.Projects, plan.ByActive)
	}
	return t._cachedActive
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
