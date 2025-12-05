package render

import (
	"github.com/SachaCR/tfb/internals/music"
	"github.com/SachaCR/tfb/internals/neck"
	"strconv"
)

func RenderScale(neck *neck.Neck, scale *music.Scale, from int, to int, mode string) string {
	if from >= to {
		return ""
	}

	strings := neck.Strings()

	renderString := ""

	if scale.Name() != "" {
		renderString = renderString + "Scale: " + scale.Root().String() + " " + scale.Name() + "\n\n"
	}

	if from < 1 {
		from = 1
	}

	if from > 1 {
		renderString = renderString + "    " + strconv.Itoa(from) + "ft\n"
	}

	// Rendering backwards to have high E string on top
	for i := len(strings) - 1; i >= 0; i-- {
		fretString := strings[i]
		renderString = renderString + fretString.Tuning().String()

		isNut := from == 1

		stringPosition := Middle

		if i == len(strings)-1 {
			stringPosition = Top
		}

		if i == 0 {
			stringPosition = Bottom
		}

		renderString = renderString + initString(stringPosition, isNut, DefaultChordStyle)

		for i := from; i <= to; i++ {
			currentFretNote := fretString.FretToNote(i)

			if scale.Contains(currentFretNote) {

				if mode == "note" {
					renderString = renderString + renderNoteSymbol(stringPosition, currentFretNote, scale, DefaultChordStyle)
					continue
				}

				renderString = renderString + renderCircleSymbol(stringPosition, currentFretNote, scale.Root(), true, DefaultChordStyle)
				continue
			}

			renderString = renderString + renderEmptyFret(stringPosition, DefaultChordStyle)
		}

		renderString = renderString + "\n"
	}

	renderString = renderString + "               •         •         •         •              ••\n"

	return renderString
}
