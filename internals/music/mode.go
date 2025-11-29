package music

import (
	"strings"
)

type Mode struct {
	name      string
	intervals []Interval
}

func (mode Mode) ToScale(root Note) *Scale {

	var scaleNotes []Note

	scaleNotes = append(scaleNotes, root)

	for _, interval := range mode.intervals {
		nextNote := root.AddTone(interval.Semitones())

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

// String returns the raw interval string.
func (i Interval) String() string {
	return string(i)
}

func (mode Mode) Intervals() []string {
	intervalStrings := []string{}

	for _, interval := range mode.intervals {
		intervalStrings = append(intervalStrings, interval.String())
	}

	return intervalStrings
}

func Ionian() *Mode {
	return NewMode("Ionian", []Interval{Interval2, Interval3, Interval4, Interval5, Interval6, Interval7})
}

func Dorian() *Mode {
	return NewMode("Dorian", []Interval{Interval2, Interval3m, Interval4, Interval5, Interval6, Interval7m})
}

func Chromatic() *Mode {
	return NewMode("Chromatic", []Interval{Interval2m, Interval2, Interval3m, Interval3, Interval4, Interval5b, Interval5, Interval6m, Interval6, Interval7m, Interval7})
}

func FindMode(modeName string) *Mode {
	switch strings.ToLower(modeName) {
	case "ionian":
		return Ionian()

	case "major":
		return Ionian()

	case "dorian":
		return Dorian()

	case "chromatic":
		return Chromatic()

	default:
		return nil
	}
}

func NewMode(name string, intervals []Interval) *Mode {
	return &Mode{
		name,
		intervals,
	}
}
