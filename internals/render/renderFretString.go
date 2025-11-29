package render

import (
	"github.com/SachaCR/tfb/internals/frets"
	"github.com/SachaCR/tfb/internals/music"
	"strconv"
	"strings"
)

func RenderFretString(fretString frets.NeckString, from int, to int, fret string, root string, stringPosition StringPosition, style ChordStyle) string {

	if from >= to {
		return ""
	}

	rootNote, isRootNoteValid := music.ParseNote(root)
	isStringRoot := fretString.Tuning().String() == rootNote.String()

	renderString := fretString.Tuning().String()

	isNut := from == 1

	switch fret {
	case "x":
		renderString = renderString + initMutedString(stringPosition, isNut, style)
	case "0":
		renderString = renderString + initOpenString(stringPosition, isNut, style, isStringRoot)
	default:
		renderString = renderString + initString(stringPosition, isNut, style)
	}

	for i := from; i <= to; i++ {

		if fret == "x" || fret == "0" {
			renderString = renderString + renderEmptyFret(stringPosition, style)
			continue
		}

		fretAsInt, err := strconv.Atoi(fret)

		if err != nil {
			renderString = renderString + renderEmptyFret(stringPosition, style)
			continue
		}

		if fretAsInt == i {
			renderString = renderString + renderCircleSymbol(stringPosition, fretString.FretToNote(i), rootNote, isRootNoteValid, style)
			continue
		}

		renderString = renderString + renderEmptyFret(stringPosition, style)
	}

	return renderString
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

		if value < min && value > 0 {
			min = value
		}

		if value > max {
			max = value
		}
	}

	if min == 99 || min == 0 || min < 3 {
		min = 1
	}

	if min == max {
		min = min - 1
		max = max + 1
	}

	return min, max
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

func red(string string) string {
	return "\033[31m" + string + "\033[0m"
}

func initMutedString(stringPosition StringPosition, isNut bool, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, isNut, style)
	return " " + string(style.mutedStringSymbol) + fretSymbol + style.stringSymbol
}

func initOpenString(stringPosition StringPosition, isNut bool, style ChordStyle, isRoot bool) string {
	fretSymbol := determineFretSymbol(stringPosition, isNut, style)

	if isRoot {
		return " " + red(style.openStringSymbol) + fretSymbol + style.stringSymbol
	} else {
		return " " + style.openStringSymbol + fretSymbol + style.stringSymbol
	}
}

func initString(stringPosition StringPosition, isNut bool, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, isNut, style)
	return "  " + fretSymbol + style.stringSymbol
}

func renderEmptyFret(stringPosition StringPosition, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, false, style)
	return style.stringSymbol + style.stringSymbol + style.stringSymbol + fretSymbol + style.stringSymbol
}

func renderNoteSymbol(stringPosition StringPosition, currentNote music.Note, scale *music.Scale, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, false, style)
	noteSymbol := currentNote.String()

	if strings.Contains(noteSymbol, "/") {
		noteSymbol = scale.ResolveEnharmonic(currentNote)
	}

	if currentNote == scale.Root() {
		noteSymbol = red(noteSymbol)
	}

	if len(currentNote.String()) == 1 {
		noteSymbol = noteSymbol + style.stringSymbol
	}

	return noteSymbol + style.stringSymbol + fretSymbol + style.stringSymbol
}

func renderCircleSymbol(stringPosition StringPosition, currentNote music.Note, rootNote music.Note, renderRoot bool, style ChordStyle) string {
	fretSymbol := determineFretSymbol(stringPosition, false, style)
	noteSymbol := style.noteSymbol

	if currentNote == rootNote && renderRoot {
		noteSymbol = red(noteSymbol)
	}

	return noteSymbol + style.stringSymbol + fretSymbol + style.stringSymbol
}
