package frets

import (
	"github.com/SachaCR/tfb/internals/music"
)

type NeckString struct {
	openStringNote music.Note
}

func (neckString NeckString) FretToNote(fretNumber int) music.Note {

	stringNoteIndex := neckString.openStringNote.ToInt()

	fretIndex := (fretNumber + stringNoteIndex) % 12

	return music.FromInt(fretIndex)
}

func (neckString NeckString) Tuning() music.Note {
	return neckString.openStringNote
}

func NewNeckString(openStringNote music.Note) NeckString {
	return NeckString{
		openStringNote: openStringNote,
	}
}
