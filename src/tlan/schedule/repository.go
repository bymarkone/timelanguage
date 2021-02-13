package schedule

type Repository struct {
	tracks []*Track
	slots  []*Slot
}

var repository Repository

func CreateRepository() {
	repository = Repository{}
}

func GetRepository() *Repository {
	return &repository
}

func (r *Repository) AddTrack(track *Track) {
	r.tracks = append(r.tracks, track)
}

func (r *Repository) ListTracks() []*Track {
	return r.tracks
}

func (r *Repository) AddSlot(slot *Slot) {
	r.slots = append(r.slots, slot)
}

func (r *Repository) ListSlots() []*Slot {
	return r.slots
}

func (r *Repository) GetSlot(name string) *Slot {
	return FindSlot(r.slots, BySlotName(name))
}

func BySlotName(name string) func(slot Slot) bool {
	return func(slot Slot) bool { return slot.Name == name }
}

func FindSlot(arr []*Slot, cond func(slot Slot) bool) *Slot {
	var result *Slot
	for i := range arr {
		if cond(*arr[i]) {
			result = arr[i]
		}
	}
	return result
}
