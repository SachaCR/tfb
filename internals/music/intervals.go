package music

type Interval int

const (
	second_minor Interval = iota
	second
	third_minor
	third
	fourth
	fourth_augmented_diminished_fifth
	fifth
	sixt_minor
	sixt
	seven_minor
	seven
)

// NoteFromString maps note names to Note constants
var IntervalFromString = map[string]Interval{

	"2m": second_minor,
	"2":  second,
	"3m": third_minor,
	"3":  third,
	"4":  fourth,
	"4+": fourth_augmented_diminished_fifth,
	"5b": fourth_augmented_diminished_fifth,
	"5":  fifth,
	"6m": sixt_minor,
	"6":  sixt,
	"7m": seven_minor,
	"7":  seven,
}
