package render

import (
	"errors"
	"github.com/SachaCR/tfb/internals/neck"
	"strconv"
	"strings"
)

func RenderChord(neck *neck.Neck, chord string, root string, chordName string) (string, error) {

	fretsToDraw := strings.Split(chord, "-")

	if len(fretsToDraw) != neck.StringCount() {
		return "", errors.New("Chord must have same string count than your instrument")
	}

	minFret, maxFret := calculateChordWidth(chord)

	renderString := ""

	if minFret > 1 {
		renderString = "     " + strconv.Itoa(minFret) + "ft"
	}

	neckStrings := neck.Strings()

	// We render backward to have the high E on top and the low E bottom
	for i := len(neckStrings) - 1; i >= 0; i-- {
		fret := fretsToDraw[i]
		fretString := neckStrings[i]

		stringPosition := Middle

		if i == len(neckStrings)-1 {
			stringPosition = Top
		}

		if i == 0 {
			stringPosition = Bottom
		}

		renderString = renderString + "\n" + RenderFretString(fretString, minFret, maxFret, fret, root, stringPosition, DefaultChordStyle)
	}

	if chordName != "" {
		renderString = renderString + "\n\t" + root + " " + chordName
	}

	return renderString, nil
}
