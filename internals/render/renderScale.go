package render

import (
	"github.com/SachaCR/neck/internals/music"
	"github.com/SachaCR/neck/internals/neck"
)

func RenderScale(neck neck.Neck, scale music.Scale, root music.Note, from int, to int) string {
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

		renderString = renderString + initString(Middle, isNut, DefaultChordStyle)

		for i := from; i <= to; i++ {

			if scale.Contains(fretString.FretToNote(i)) {
				renderString = renderString + renderNoteSymbol(Middle, fretString.FretToNote(i), scale.Root(), true, DefaultChordStyle)
				continue
			}

			renderString = renderString + renderEmptyFret(Middle, DefaultChordStyle)

		}

		renderString = renderString + "\n"
	}

	return renderString
}
