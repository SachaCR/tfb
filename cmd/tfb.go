package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/music"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&instrument, "inst", "i", "G", "Set the instrument type, G for guitar, B for Bass, U for Ukulele")
	rootCmd.PersistentFlags().StringVarP(&root, "root", "r", "", "Set the root of your chord")
}

var rootCmd = &cobra.Command{
	Use:   "tfb",
	Short: "tfb is a tools for guitarists that display neck diagrams",
	Long:  `tfb allows to display guitar neck diagrams. It also support bass and ukulele and potentially any instruments with strings and frets.`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}
		modeName := args[0]

		mode := music.FindMode(modeName)

		if mode == nil {
			fmt.Println("Unknow Mode")
		} else {
			fmt.Println(mode.Intervals())
		}

		neck := neck.New(instrument)

		scale, err := music.ParseScale("C-D-E-F-G-A-B", "Major")

		if err != nil {
			panic(err)
		}

		fmt.Println(render.RenderScale(neck, scale, 1, 12, "note"))
	},
}

func init() {
	rootCmd.AddCommand(chordCmd)
	rootCmd.AddCommand(scaleCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
