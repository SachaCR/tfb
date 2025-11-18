package neck

import (
	"neck/internals/frets"
	"neck/internals/music"
)

type Neck struct {
	instrument string
	strings    []frets.FretString
}

func (neck *Neck) AddString(note music.Note) {
	neck.strings = append(neck.strings, frets.NewFretString(note))
}

func (neck Neck) Tuning() string {
	tuning := ""

	for _, fretString := range neck.strings {
		tuning = tuning + fretString.Tuning()
	}

	return tuning
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
