package render

import (
	"errors"
	"github.com/SachaCR/neck/internals/frets"
	"github.com/SachaCR/neck/internals/music"
	"github.com/SachaCR/neck/internals/neck"
	"strconv"
	"strings"
)

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
	max := 3

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

	return min, max
}

func RenderChord(neck neck.Neck, chord string, root string, chordName string) (string, error) {

	fretsToDraw := strings.Split(chord, "-")

	if len(fretsToDraw) != neck.StringCount() {
		return "", errors.New("Chord must have same string count than your instrument")
	}

	minFret, maxFret := calculateChordWidth(chord)

	renderString := ""

	noteList := neck.Tuning()

	if minFret > 1 {
		renderString = "     " + strconv.Itoa(minFret) + "th"
	}

	// We render backward to have the high E on top and the low E bottom
	for i := len(fretsToDraw) - 1; i >= 0; i-- {
		fret := fretsToDraw[i]
		fretString := frets.NewFretString(noteList[i])

		stringPosition := Middle

		if i == len(fretsToDraw)-1 {
			stringPosition = Top
		}

		if i == 0 {
			stringPosition = Bottom
		}

		renderString = renderString + "\n" + RenderFretString(fretString, minFret, maxFret, fret, root, stringPosition)
	}

	if chordName != "" {
		renderString = renderString + "\n\t" + root + " " + chordName
	}

	return renderString, nil
}

func determineFretSymbol(stringPosition StringPosition, isNut bool, style ChordStyle) string {
	if isNut {
		switch stringPosition {
		case Top:
			return style.topFirstFretSymbol
		case Bottom:
			return style.bottomFirstFretSymbol
		default:
			return style.middleFirstFretSymbol
		}
	}

	switch stringPosition {
	case Top:
		return style.topFretSymbol
	case Bottom:
		return style.bottomFretSymbol
	default:
		return style.fretSymbol
	}
}

func initMutedString(stringPosition StringPosition, isNut bool, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, isNut, DefaultChordStyle)
	return " " + string(style.mutedStringSymbol) + fretSymbol + style.stringSymbol
}

func initOpenString(stringPosition StringPosition, isNut bool, style ChordStyle, isRoot bool) string {
	fretSymbol := determineFretSymbol(stringPosition, isNut, DefaultChordStyle)

	if isRoot {
		// Display the 0 in red when it's the chord's root
		return " " + "\033[31m" + style.openStringSymbol + "\033[0m" + fretSymbol + style.stringSymbol
	} else {
		return " " + style.openStringSymbol + fretSymbol + style.stringSymbol
	}
}

func initString(stringPosition StringPosition, isNut bool, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, isNut, DefaultChordStyle)
	return "  " + fretSymbol + style.stringSymbol
}

func renderEmptyFret(stringPosition StringPosition, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, false, DefaultChordStyle)
	return style.stringSymbol + style.stringSymbol + style.stringSymbol + fretSymbol + style.stringSymbol
}

func renderNoteSymbol(stringPosition StringPosition, currentNote music.Note, rootNote music.Note, renderRoot bool, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, false, DefaultChordStyle)

	noteSymbol := "⬤ "

	if currentNote == rootNote && renderRoot {
		noteSymbol = "🔴"
	}

	return noteSymbol + style.stringSymbol + fretSymbol + style.stringSymbol
}

func RenderFretString(fretString frets.FretString, from int, to int, fret string, root string, stringPosition StringPosition) string {

	if from >= to {
		return ""
	}

	rootNote, isRootNoteValid := music.ParseNote(root)
	isStringRoot := fretString.Tuning().String() == rootNote.String()

	renderString := fretString.Tuning().String()

	isNut := from == 1

	if fret == "x" {
		renderString = renderString + initMutedString(stringPosition, isNut, DefaultChordStyle)
	} else if fret == "0" {
		renderString = renderString + initOpenString(stringPosition, isNut, DefaultChordStyle, isStringRoot)
	} else {
		renderString = renderString + initString(stringPosition, isNut, DefaultChordStyle)
	}

	for i := from; i <= to; i++ {

		if fret == "x" || fret == "0" {
			renderString = renderString + renderEmptyFret(stringPosition, DefaultChordStyle)
			continue
		}

		fretAsInt, err := strconv.Atoi(fret)

		if err != nil {
			renderString = renderString + renderEmptyFret(stringPosition, DefaultChordStyle)
			continue
		}

		if fretAsInt == i {
			renderString = renderString + renderNoteSymbol(stringPosition, fretString.FretToNote(i), rootNote, isRootNoteValid, DefaultChordStyle)
			continue
		}

		renderString = renderString + renderEmptyFret(stringPosition, DefaultChordStyle)
	}

	return renderString
}
