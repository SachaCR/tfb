package music

import (
	"strings"
)

type Mode struct {
	name      string
	intervals []Interval
}

func (mode Mode) Name() string {
	return mode.name
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

var ModeMap = map[string]*Mode{}

func init() {
	ModeMap["ionian"] = NewMode("Ionian", []Interval{Interval2, Interval3, Interval4, Interval5, Interval6, Interval7})
	ModeMap["dorian"] = NewMode("Dorian", []Interval{Interval2, Interval3m, Interval4, Interval5, Interval6, Interval7m})
	ModeMap["phrygian"] = NewMode("Phrygian", []Interval{Interval2m, Interval3m, Interval4, Interval5, Interval6m, Interval7m})
	ModeMap["lydian"] = NewMode("Lydian", []Interval{Interval2, Interval3, Interval4p, Interval5, Interval6, Interval7})
	ModeMap["mixolydian"] = NewMode("Mixolydian", []Interval{Interval2, Interval3, Interval4, Interval5, Interval6, Interval7})
	ModeMap["aeolian"] = NewMode("Aeolian", []Interval{Interval2, Interval3m, Interval4, Interval5, Interval6m, Interval7m})
	ModeMap["locrian"] = NewMode("Locrian", []Interval{Interval2m, Interval3m, Interval4, Interval5b, Interval6m, Interval7m})
	ModeMap["chromatic"] = NewMode("Chromatic", []Interval{Interval2m, Interval2, Interval3m, Interval3, Interval4, Interval5b, Interval5, Interval6m, Interval6, Interval7m, Interval7})
}

func FindMode(modeName string) *Mode {
	mode := ModeMap[strings.ToLower(modeName)]
	return mode
}

func NewMode(name string, intervals []Interval) *Mode {
	return &Mode{
		name,
		intervals,
	}
}
