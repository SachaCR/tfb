package music

type Interval string

const (
	Interval2m Interval = "2m" // minor second
	Interval2  Interval = "2"  // major second
	Interval3m Interval = "3m" // minor third
	Interval3  Interval = "3"  // major third
	Interval4  Interval = "4"  // perfect fourth
	Interval4p Interval = "4+" // augmented fourth (tritone)
	Interval5b Interval = "5b" // diminished fifth (tritone)
	Interval5  Interval = "5"  // perfect fifth
	Interval5p Interval = "5+" // augmented fifth
	Interval6m Interval = "6m" // minor sixth
	Interval6  Interval = "6"  // major sixth
	Interval7m Interval = "7m" // minor seventh
	Interval7  Interval = "7"  // major seventh
)

var intervals = []Interval{
	Interval2m,
	Interval2,
	Interval3m,
	Interval3,
	Interval4,
	Interval4p,
	Interval5b,
	Interval5,
	Interval5p,
	Interval6m,
	Interval6,
	Interval7m,
	Interval7,
}

// Semitone mapping for each interval
var intervalSemitones = map[Interval]int{
	Interval2m: 1,
	Interval2:  2,
	Interval3m: 3,
	Interval3:  4,
	Interval4:  5,
	Interval4p: 6,
	Interval5b: 6,
	Interval5:  7,
	Interval5p: 8,
	Interval6m: 8,
	Interval6:  9,
	Interval7m: 10,
	Interval7:  11,
}

// Semitones returns the number of semitones in the interval.
func (i Interval) Semitones() int {
	return intervalSemitones[i]
}

func IntervalFromInt(semitones int) Interval {
	n := len(intervals)
	// Ensure positive modulo
	idx := ((semitones % n) + n) % n
	return intervals[idx]
}
