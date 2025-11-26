package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/spf13/cobra"
)

var chordRoot string
var chordName string

func init() {
	chordCmd.PersistentFlags().StringVarP(&chordName, "name", "n", "", "Give a chord name like Major7 or m7b5")
	chordCmd.PersistentFlags().StringVarP(&chordRoot, "root", "r", "", "Set the root of your chord")
}

var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Print chords",
	Long:  `Print chords that you pass in argument. Example: 0-2-2-0-0-0`,

	Run: func(cmd *cobra.Command, args []string) {

		chord := args[0]
		neck := neck.GuitarNeck()

		chordAsString, err := render.RenderChord(neck, chord, chordRoot, chordName)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(chordAsString)
	},
}
