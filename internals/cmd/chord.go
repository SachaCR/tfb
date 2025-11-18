package cmd

import (
	"fmt"
	"github.com/SachaCR/neck/internals/music"
	"github.com/SachaCR/neck/internals/neck"
	"github.com/SachaCR/neck/internals/render"
	"github.com/spf13/cobra"
)

var name string
var root string

func init() {
	chordCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Give a chord name like Major7 or m7b5")
	chordCmd.PersistentFlags().StringVarP(&root, "root", "r", "", "Set the root of your chord")
}

var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Print chords",
	Long:  `Print chords that you pass in argument. Example: 0-2-2-0-0-0`,

	Run: func(cmd *cobra.Command, args []string) {
		rootNote, ok := music.ParseNote(root)

		if !ok {
			panic("Inalid root note")
		}

		chord := args[0]
		neck := neck.GuitarNeck()

		chordAsString, err := render.RenderChord(neck, chord, rootNote, "Minor")

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(chordAsString)
	},
}
