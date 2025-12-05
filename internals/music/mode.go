package music

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
	major := []Interval{Interval2, Interval3, Interval4, Interval5, Interval6, Interval7}
	minor := []Interval{Interval2, Interval3m, Interval4, Interval5, Interval6m, Interval7m}

	// Basic Modes
	ModeMap["Major"] = NewMode("Major", major)
	ModeMap["Minor"] = NewMode("Minor", minor)

	// Musical Modes
	ModeMap["Ionian"] = NewMode("Ionian", major)
	ModeMap["Dorian"] = NewMode("Dorian", []Interval{Interval2, Interval3m, Interval4, Interval5, Interval6, Interval7m})
	ModeMap["Phrygian"] = NewMode("Phrygian", []Interval{Interval2m, Interval3m, Interval4, Interval5, Interval6m, Interval7m})
	ModeMap["Lydian"] = NewMode("Lydian", []Interval{Interval2, Interval3, Interval4p, Interval5, Interval6, Interval7})
	ModeMap["Mixolydian"] = NewMode("Mixolydian", []Interval{Interval2, Interval3, Interval4, Interval5, Interval6, Interval7m})
	ModeMap["Aeolian"] = NewMode("Aeolian", minor)
	ModeMap["Locrian"] = NewMode("Locrian", []Interval{Interval2m, Interval3m, Interval4, Interval5b, Interval6m, Interval7m})

	// Pentatonic Modes
	ModeMap["Minor Pentatonic"] = NewMode("Minor Pentatonic", []Interval{Interval3m, Interval4, Interval5, Interval7m})
	ModeMap["Major Pentatonic"] = NewMode("Major Pentatonic", []Interval{Interval2, Interval3, Interval5, Interval6})
	ModeMap["Egyptian"] = NewMode("Egyptian", []Interval{Interval2, Interval4, Interval5, Interval7m})
	ModeMap["Quang ming"] = NewMode("Quand Ming", []Interval{Interval3m, Interval4, Interval5p, Interval7m})
	ModeMap["Ritusen"] = NewMode("Ritusen", []Interval{Interval2, Interval4, Interval5, Interval6})

	// ModeMap["chromatic"] = NewMode("Chromatic", []Interval{Interval2m, Interval2, Interval3m, Interval3, Interval4, Interval5b, Interval5, Interval6m, Interval6, Interval7m, Interval7})
}

func FindMode(modeName string) *Mode {
	mode, ok := ModeMap[modeName]

	if !ok {
		// This is better than crashing
		mode = ModeMap["Major"]
	}

	return mode
}

func NewMode(name string, intervals []Interval) *Mode {
	return &Mode{
		name,
		intervals,
	}
}
