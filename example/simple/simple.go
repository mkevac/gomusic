package main

import (
	"fmt"
	"github.com/mkevac/gomusic"
)

func main() {

	note1 := gomusic.Note{}

	mfile := gomusic.NewMIDIFile(1, 60)
	mfile.AddNoteSeq(0, 0, []interface{}{note1, note1})
	err := mfile.Write("test.mid")
	if err != nil {
		fmt.Println(err)
	}
}
