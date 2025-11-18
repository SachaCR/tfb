package render

import (
	"errors"
	"neck/internals/frets"
	"neck/internals/music"
	"neck/internals/neck"
	"strconv"
	"strings"
)

// Rendering example

// fmt.Println("E x┬─⬤ ─┬────┬────┬")
// fmt.Println("B  ┼─⬤ ─┼────┼────┼")
// fmt.Println("G  ┼────┼─⬤ ─┼────┼")
// fmt.Println("D  ┼────┼────┼─⬤ ─┼")
// fmt.Println("A  ┼─⬤ ─┼─Ab─┼─⬤ ─┼")
// fmt.Println("E 0┴─🔴─┴────┴────┴")

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
func calculateChordWidth(chord string) (int, int) {

	frets := strings.Split(chord, "-")
	min := 99 // TODO Fix this
	max := 0

	for _, fret := range frets {

		value, err := strconv.Atoi(fret)
		if err != nil {
			continue
		}

		if value < min && value >= 0 {
			min = value
		}

		if value > max {
			max = value
		}
	}

	if min == 99 || min == 0 {
		min = 1
	}

	if max == 0 || max < 3 {
		max = 3
	}

	return min, max
}

func RenderChord(neck neck.Neck, chord string, root music.Note) (string, error) {
	fretsToDraw := strings.Split(chord, "-")

	if len(fretsToDraw) != neck.StringCount() {
		return "", errors.New("Chord must have same string count than your instrument")
	}

	minFret, maxFret := calculateChordWidth(chord)

	noteList := neck.Tuning()
	renderString := ""

	for i := len(fretsToDraw); i > 0; i-- {
		fret := fretsToDraw[i-1]

		// To render from high to bass: Start by the last string to reverse string order when rendering.
		stringIndex := len(noteList) - i

		fretString := frets.NewFretString(noteList[stringIndex])
		renderString = renderString + "\n" + RenderFretString(fretString, minFret, maxFret, fret, root)
	}

	return renderString, nil
}

func RenderFretString(fretString frets.FretString, from int, to int, fret string, root music.Note) string {

	if from >= to {
		return ""
	}

	renderString := fretString.Tuning().String()

	if fret == "x" || fret == "X" {
		renderString = renderString + " x┼─"
	} else if fret == "0" || fret == "O" {
		renderString = renderString + " 0┼─"
	} else {
		renderString = renderString + "  ┼─"
	}

	for i := from; i <= to; i++ {
		fretSymbol := "──"

		if fret == "x" || fret == "0" {
			renderString = renderString + fretSymbol + "─┼─"
			continue
		}

		fretAsInt, err := strconv.Atoi(fret)

		if err != nil {
			renderString = renderString + fretSymbol + "─┼─"
			continue
		}

		if fretAsInt == i {
			if fretString.FretToNote(i) == root {
				fretSymbol = "🔴"
			} else {
				fretSymbol = "⬤ "
			}
		}

		renderString = renderString + fretSymbol + "─┼─"

	}

	return renderString
}
