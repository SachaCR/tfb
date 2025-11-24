package cmd

import (
	"fmt"
	"github.com/SachaCR/neck/internals/music"
	"github.com/SachaCR/neck/internals/neck"
	"github.com/SachaCR/neck/internals/render"
	"github.com/spf13/cobra"
	"strings"
)

var scaleRoot string
var scaleName string
var from int
var to int

func init() {
	scaleCmd.PersistentFlags().StringVarP(&scaleName, "name", "n", "", "Set the scale name. Example: C Major")
	scaleCmd.PersistentFlags().StringVarP(&scaleRoot, "root", "r", "", "Set the root note of your scale. By default it takes the first note of the scale.")
	scaleCmd.PersistentFlags().IntVarP(&from, "from", "f", 1, "Render from that fret number")
	scaleCmd.PersistentFlags().IntVarP(&to, "to", "t", 12, "Last fret number to render")
}

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Print scale",
	Long:  `Print scale that you pass in argument. Example: C-D-E-F-G-A-B`,

	Run: func(cmd *cobra.Command, args []string) {

		scaleInput := args[0]
		stringArray := strings.Split(scaleInput, "-")
		var noteSlice []music.Note

		for _, string := range stringArray {
			note, ok := music.NoteFromString[string]

			if !ok {
				panic("invalid note")
			}

			noteSlice = append(noteSlice, note)
		}

		var scale music.Scale
		rootNote := noteSlice[0]

		if scaleRoot != "" {
			note, ok := music.ParseNote(scaleRoot)

			if !ok {
				panic("Invalid Root Note")
			}

			rootNote = note
		}

		scale = music.NewScale(noteSlice, rootNote)

		neck := neck.GuitarNeck()

		if scaleName != "" {
			fmt.Println("Scale:", rootNote.String(), scaleName)
		}

		if from < 1 {
			from = 1
		}

		if from > 1 {
			fmt.Println("    ", from, "ft")
		}

		scaleAsString := render.RenderScale(neck, scale, scale.Root(), from, to)

		fmt.Print(scaleAsString)
		fmt.Println("               .         .         .         .              ..")
	},
}
