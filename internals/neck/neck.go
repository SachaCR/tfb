package neck

import (
	"github.com/SachaCR/tfb/internals/frets"
	"github.com/SachaCR/tfb/internals/music"
)

type Neck struct {
	instrument string
	strings    []frets.FretString
}

func (neck *Neck) AddString(note music.Note) {
	neck.strings = append(neck.strings, frets.NewFretString(note))
}

func (neck Neck) Insrtrument() string {
	return neck.instrument
}

func (neck Neck) Strings() []frets.FretString {
	return neck.strings
}

func (neck Neck) Tuning() []music.Note {
	var tuning []music.Note

	for _, fretString := range neck.strings {
		tuning = append(tuning, fretString.Tuning())
	}

	return tuning
}

func (neck Neck) StringCount() int {
	return len(neck.strings)
}

func GuitarNeck() Neck {
	var fretStrings []frets.FretString

	neck := Neck{
		instrument: "Guitare",
		strings:    fretStrings,
	}

	neck.AddString(music.E)
	neck.AddString(music.A)
	neck.AddString(music.D)
	neck.AddString(music.G)
	neck.AddString(music.B)
	neck.AddString(music.E)

	return neck
}
