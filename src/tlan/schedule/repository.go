package schedule

var tracks []*Track
var slots []*Slot

func AddTrack(track *Track) {
	tracks = append(tracks, track)
}

func ListTracks() []*Track {
	return tracks
}

func AddSlot(slot *Slot) {
	slots = append(slots, slot)
}

func ListSlots() []*Slot {
	return slots
}

func GetSlot(name string) *Slot {
	return FindSlot(slots, BySlotName(name))
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

func Clean() {
	tracks = []*Track{}
}
