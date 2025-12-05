package neck

import (
	"github.com/SachaCR/tfb/internals/frets"
	"github.com/SachaCR/tfb/internals/music"
)

type Neck struct {
	instrument string
	strings    []frets.NeckString
}

func (neck *Neck) AddString(note music.Note) {
	neck.strings = append(neck.strings, frets.NewNeckString(note))
}

func (neck Neck) Instrument() string {
	return neck.instrument
}

func (neck Neck) Strings() []frets.NeckString {
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

func New(instrument string) *Neck {
	switch instrument {
	case "B":
		return BassNeck()

	case "U":
		return UkuleleNeck()

	default:
		return GuitarNeck()
	}
}
func UkuleleNeck() *Neck {

	var fretStrings []frets.NeckString

	neck := Neck{
		instrument: "Ukulele",
		strings:    fretStrings,
	}

	neck.AddString(music.G)
	neck.AddString(music.C)
	neck.AddString(music.E)
	neck.AddString(music.A)

	return &neck
}
func BassNeck() *Neck {

	var fretStrings []frets.NeckString

	neck := Neck{
		instrument: "Bass",
		strings:    fretStrings,
	}

	neck.AddString(music.E)
	neck.AddString(music.A)
	neck.AddString(music.D)
	neck.AddString(music.G)

	return &neck

}
func GuitarNeck() *Neck {
	var fretStrings []frets.NeckString

	neck := Neck{
		instrument: "Guitar",
		strings:    fretStrings,
	}

	neck.AddString(music.E)
	neck.AddString(music.A)
	neck.AddString(music.D)
	neck.AddString(music.G)
	neck.AddString(music.B)
	neck.AddString(music.E)

	return &neck
}
