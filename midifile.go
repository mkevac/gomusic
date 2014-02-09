package gomusic

import (
	"errors"
	"github.com/mkevac/gomidicreator"
	"os"
)

type MIDIFile struct {
	tracks     uint16
	tempo      int
	instrument int
	MidiData   gomidicreator.MIDIFile
}

func NewMIDIFile(tracks uint16, tempo int) MIDIFile {
	file := MIDIFile{
		tracks:   tracks,
		tempo:    tempo,
		MidiData: gomidicreator.NewMIDIFile(tracks),
	}

	var i uint16
	for i = 0; i < tracks; i++ {
		file.MidiData.AddTrackName(i, 0, "Track #"+string(i))
		file.MidiData.SetTempo(i, 0, tempo)
		file.MidiData.SetProgramChange(i, 0, 0)
	}

	return file
}

func (midi MIDIFile) AddNoteSeq(track uint16, time int, seq []interface{}) (int, error) {
	if track > midi.tracks {
		return time, errors.New("wrong track")
	}

	for _, seqItem := range seq {
		switch note := seqItem.(type) {
		case Note:
			midi.MidiData.AddNote(track, note.Value, time, note.midiDuration, note.Volume)
			time += note.midiDuration
		case Rest:
			time += note.midiDuration
		default:
			return time, errors.New("wrong type in list")
		}
	}

	return time, nil
}

func (midi MIDIFile) Write(filename string) error {

	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	err = midi.MidiData.WriteFile(fd)
	if err != nil {
		return err
	}

	return nil
}
