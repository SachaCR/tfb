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

func init() {
	scaleCmd.PersistentFlags().StringVarP(&scaleName, "name", "n", "", "Set the scale name. Example: C Major")
	scaleCmd.PersistentFlags().StringVarP(&scaleRoot, "root", "r", "", "Set the root note of your scale")
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
		scaleAsString := render.RenderScale(neck, scale, scale.Root(), 1, 12)

		fmt.Print(scaleAsString)
		fmt.Println("               .         .         .         .              ..")
	},
}
