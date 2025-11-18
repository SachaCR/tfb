package render

import (
	"neck/internals/frets"
)

// fmt.Println("┬─⬤ ─┬────┬────┬")
// fmt.Println("┼─⬤ ─┼────┼────┼")
// fmt.Println("┼────┼─⬤ ─┼────┼")
// fmt.Println("┼────┼────┼─⬤ ─┼")
// fmt.Println("┼─⬤ ─┼─Ab─┼─⬤ ─┼")
// fmt.Println("┴─⬤ ─┴────┴────┴")

func contains(slice []int, search int) bool {
	found := false

	for _, n := range slice {
		if n == search {
			found = true
			break
		}
	}
	return found
}

func RenderFretString(fretString frets.FretString, from int, to int, mark []int) string {

	if from >= to {
		return ""
	}

	renderString := fretString.Tuning() + "┼─"
	for i := from; i <= to; i++ {
		fretSymbol := "──"
		if contains(mark, i) {
			fretSymbol = "⬤ "
		}

		renderString = renderString + fretSymbol + "─┼─"

	}

	return renderString
}
