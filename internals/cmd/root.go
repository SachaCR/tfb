package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "neck",
	Short: "Neck is a tools for guitarists that display neck diagrams",
	Long:  `Neck allows to display guitar neck diagrams. It also support bass and ukulele and potentially any instruments with strings and frets.`,
	Run: func(cmd *cobra.Command, args []string) {
		println("neck")
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
