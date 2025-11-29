package music

// Note represents a musical note
type Note int

const (
	A Note = iota
	ASharp_Bb
	B
	C
	CSharp_Db
	D
	DSharp_Eb
	E
	F
	FSharp_Gb
	G
	GSharp_Ab
)

// NoteFromString maps note names to Note constants
var NoteFromString = map[string]Note{
	"A":  A,
	"A#": ASharp_Bb,
	"Bb": ASharp_Bb,
	"B":  B,
	"C":  C,
	"C#": CSharp_Db,
	"Db": CSharp_Db,
	"D":  D,
	"D#": DSharp_Eb,
	"Eb": DSharp_Eb,
	"E":  E,
	"F":  F,
	"F#": FSharp_Gb,
	"Gb": FSharp_Gb,
	"G":  G,
	"G#": GSharp_Ab,
	"Ab": GSharp_Ab,
}

// ParseNote converts a string to a Note, returns ok=false if invalid
func ParseNote(s string) (Note, bool) {
	n, ok := NoteFromString[s]
	return n, ok
}

// FromInt converts an integer to a Note
func FromInt(i int) Note {
	if i < 0 {
		i = (i%12 + 12) % 12 // handle negative numbers
	}

	return Note(i % 12)
}

// ToInt returns the integer representation of the note (0-11)
func (n Note) ToInt() int {
	return int(n % 12)
}

func (n Note) AddTone(tone int) Note {

	newNote := FromInt(n.ToInt() + tone)

	return newNote

}

// String returns the note name
func (n Note) String() string {
	return [...]string{
		"A",
		"A#/Bb",
		"B",
		"C",
		"C#/Db",
		"D",
		"D#/Eb",
		"E",
		"F",
		"F#/Gb",
		"G",
		"G#/Ab",
	}[n]
}
