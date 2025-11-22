package frets

import (
	"github.com/SachaCR/neck/internals/music"
)

type FretString struct {
	openStringNote music.Note
}

func (fretString FretString) FretToNote(fretNumber int) music.Note {

	stringNoteIndex := fretString.openStringNote.ToInt()

	fretIndex := (fretNumber + stringNoteIndex) % 12

	return music.FromInt(fretIndex)
}

func (fretString FretString) Tuning() music.Note {
	return fretString.openStringNote
}

func NewFretString(openStringNote music.Note) FretString {
	return FretString{
		openStringNote: openStringNote,
	}
}
