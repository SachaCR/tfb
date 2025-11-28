package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/music"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tfb",
	Short: "tfb is a tools for guitarists that display neck diagrams",
	Long:  `tfb allows to display guitar neck diagrams. It also support bass and ukulele and potentially any instruments with strings and frets.`,
	Run: func(cmd *cobra.Command, args []string) {
		neck := neck.GuitarNeck()
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
