package gomusic

const (
	Quarter = 0.25
	Half    = 0.5
	Whole   = 1.0
)

type Note struct {
	Value        int
	Octave       int
	Duration     float32
	Volume       int
	midiNumber   int
	midiDuration int
}

type Rest struct {
	Duration     float32
	midiDuration int
}
