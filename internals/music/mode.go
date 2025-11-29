package music

type Mode struct {
	name      string
	intervals []int
}

func (mode Mode) ToScale(root Note) *Scale {

	var scaleNotes []Note

	scaleNotes = append(scaleNotes, root)

	for _, tone := range mode.intervals {
		previousNote := scaleNotes[len(scaleNotes)-1]
		nextNote := previousNote.AddTone(tone)

		if nextNote == root || len(scaleNotes) == 12 {
			break
		}

		scaleNotes = append(scaleNotes, nextNote)
	}

	scale, err := NewScale(scaleNotes, root, mode.name, []string{})
	if err != nil {
		panic("Invalid Scale")
	}

	return scale
}

func Ionian() *Mode {
	return NewMode("Ionian", []int{2, 2, 1, 2, 2, 2, 1})
}

func Dorian() *Mode {
	return NewMode("Dorian", []int{2, 1, 2, 2, 2, 1, 2})
}

func Chromatic() *Mode {
	return NewMode("Chromatic", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
}

func NewMode(name string, intervals []int) *Mode {
	return &Mode{
		name,
		intervals,
	}
}
