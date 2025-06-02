package player

type AudioFile struct {
	name     string
	path     string
	duration float64
}

func (a *AudioFile) getName() string {
	return a.name
}

func (a *AudioFile) getPath() string {
	return a.path
}

func (a *AudioFile) getDuration() float64 {
	return a.duration
}
