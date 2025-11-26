package render

import (
	"github.com/SachaCR/tfb/internals/music"
	"github.com/SachaCR/tfb/internals/neck"
)

func RenderScale(neck neck.Neck, scale music.Scale, root music.Note, from int, to int, mode string) string {
	if from >= to {
		return ""
	}

	strings := neck.Strings()

	renderString := ""

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

			if scale.Contains(fretString.FretToNote(i)) {
				if mode == "note" {
					renderString = renderString + renderNoteSymbol(stringPosition, fretString.FretToNote(i), scale.Root(), DefaultChordStyle)
					continue
				}

				renderString = renderString + renderCircleSymbol(stringPosition, fretString.FretToNote(i), scale.Root(), true, DefaultChordStyle)
				continue
			}

			renderString = renderString + renderEmptyFret(stringPosition, DefaultChordStyle)
		}

		renderString = renderString + "\n"
	}

	return renderString
}
