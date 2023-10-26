package schedule

import (
	plan "github.com/bymarkone/timelanguage/internal/planning"
	"github.com/bymarkone/timelanguage/internal/utils"
)

type Track struct {
	Slot                 *Slot
	Name                 string
	Projects             []*plan.Project
	_cachedActive        []*plan.Project
	_cachedFlattenActive []*plan.Project
}

func (t *Track) ActiveProjects() []*plan.Project {
	if len(t._cachedActive) == 0 {
		t._cachedActive = plan.FilterProjects(t.Projects, plan.ProjectActive)
	}
	return t._cachedActive
}

func (t *Track) FlattenActiveProjects() []*plan.Project {
	if len(t._cachedFlattenActive) == 0 {
		t._cachedFlattenActive = plan.FilterProjects(plan.FlattenProjects(t.Projects), plan.ProjectActive)
	}
	return t._cachedFlattenActive
}

type Slot struct {
	Name                      string
	Period                    utils.Period
	Tracks                    []*Track
	_cachedFlattenActiveItems []string
}

func (s *Slot) FlattenActiveItems(flattener func(arr []*Track) []string) []string {
	if len(s._cachedFlattenActiveItems) == 0 {
		s._cachedFlattenActiveItems = flattener(s.Tracks)
	}
	return s._cachedFlattenActiveItems
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
