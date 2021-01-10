package schedule

var tracks []*Track

func AddTrack(track Track) {
	tracks = append(tracks, &track)
}

func ListTracks() []*Track {
	return tracks
}

func Clean() {
	tracks = []*Track{}
}
