package render

// first: "┬────",
// normal: "┼────",
// last: "┴────",

// fmt.Println("E x┬─⬤ ─┬────┬────┬")
// fmt.Println("B  ┼─⬤ ─┼────┼────┼")
// fmt.Println("G  ┼────┼─⬤ ─┼────┼")
// fmt.Println("D  ┼────┼────┼─⬤ ─┼")
// fmt.Println("A  ┼─⬤ ─┼─Ab─┼─⬤ ─┼")
// fmt.Println("E 0┴─🔴─┴────┴────┴")

type ChordStyle struct {
	topFretSymbol         string
	fretSymbol            string
	bottomFretSymbol      string
	stringSymbol          string
	rootSymbol            string
	noteSymbol            string
	openStringSymbol      string
	mutedStringSymbol     string
	topFirstFretSymbol    string
	middleFirstFretSymbol string
	bottomFirstFretSymbol string
}

var DefaultChordStyle = ChordStyle{
	topFretSymbol:         "┬",
	fretSymbol:            "┼",
	bottomFretSymbol:      "┴",
	stringSymbol:          "─",
	rootSymbol:            "🔴",
	noteSymbol:            "⬤ ",
	openStringSymbol:      "0",
	mutedStringSymbol:     "x",
	topFirstFretSymbol:    "╓",
	middleFirstFretSymbol: "╟",
	bottomFirstFretSymbol: "╙",
}
