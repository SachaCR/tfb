package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/spf13/cobra"
)

func init() {
	chordCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Give a chord name like Major7 or m7b5")
	chordCmd.PersistentFlags().StringVarP(&root, "root", "r", "", "Set the root of your chord")
	chordCmd.PersistentFlags().StringVarP(&instrument, "inst", "i", "G", "Set the instrument type, G for Guitar, B for Bass, U for Ukulele")
	chordCmd.PersistentFlags().StringVarP(&tuning, "tuning", "t", "E-A-D-G-B-E", "Set a custom instrument tuning")
}

var chordCmd = &cobra.Command{
	Use:   "chord",
	Short: "Print chords",
	Long:  `Print chords that you pass in argument. Example: 0-2-2-0-0-0`,

	Run: func(cmd *cobra.Command, args []string) {

		chord := args[0]

		var instrumentNeck *neck.Neck
		var err error

		instrumentNeck = neck.New(instrument)
		if tuning != "" {
			instrumentNeck, err = neck.NewCustom("Custom instrument", tuning)
			if err != nil {
				panic(err)
			}
		}

		chordAsString, err := render.RenderChord(instrumentNeck, chord, root, name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(chordAsString)
	},
}
