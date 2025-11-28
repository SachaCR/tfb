package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/spf13/cobra"
)

var chordRoot string
var chordName string
var instrumentType string

func init() {
	chordCmd.PersistentFlags().StringVarP(&chordName, "name", "n", "", "Give a chord name like Major7 or m7b5")
	chordCmd.PersistentFlags().StringVarP(&chordRoot, "root", "r", "", "Set the root of your chord")
	chordCmd.PersistentFlags().StringVarP(&instrumentType, "inst", "i", "G", "Set the instrument type, G for Guitar, B for Bass, U for Ukulele")
}

var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Print chords",
	Long:  `Print chords that you pass in argument. Example: 0-2-2-0-0-0`,

	Run: func(cmd *cobra.Command, args []string) {

		chord := args[0]
		neck := neck.New(instrumentType)

		chordAsString, err := render.RenderChord(neck, chord, chordRoot, chordName)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(chordAsString)
	},
}
