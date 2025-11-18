package render

import (
	"errors"
	"github.com/SachaCR/neck/internals/frets"
	"github.com/SachaCR/neck/internals/music"
	"github.com/SachaCR/neck/internals/neck"
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

func RenderChord(neck neck.Neck, chord string, root string, chordName string) (string, error) {

	rootNote, _ := music.ParseNote(root)

	// if !ok {
	// 	rootNote = ""
	// }

	fretsToDraw := strings.Split(chord, "-")

	if len(fretsToDraw) != neck.StringCount() {
		return "", errors.New("Chord must have same string count than your instrument")
	}

	minFret, maxFret := calculateChordWidth(chord)

	noteList := neck.Tuning()

	renderString := "     " + strconv.Itoa(minFret) + "th"

	// We render backward to have the high E on top and the low E bottom
	for i := len(fretsToDraw); i > 0; i-- {
		fret := fretsToDraw[i-1]
		fretString := frets.NewFretString(noteList[i-1])
		renderString = renderString + "\n" + RenderFretString(fretString, minFret, maxFret, fret, rootNote)
	}

	if chordName != "" {
		renderString = renderString + "\n\t" + rootNote.String() + " " + chordName
	}

	return renderString, nil
}

func RenderFretString(fretString frets.FretString, from int, to int, fret string, root music.Note) string {

	if from >= to {
		return ""
	}

	renderString := fretString.Tuning().String()

	if fret == "x" {
		renderString = renderString + " x┼─"
	} else if fret == "0" {
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
