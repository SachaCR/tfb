package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/music"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/spf13/cobra"
)

func init() {
	scaleCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Set the scale name. Example: Major")
	scaleCmd.PersistentFlags().IntVarP(&from, "from", "f", 1, "Render from that fret number")
	scaleCmd.PersistentFlags().IntVarP(&to, "to", "t", 12, "Last fret number to render")
	scaleCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "circle", "Set the display mode to `circle` or `note`. Default to `circle`")
	scaleCmd.PersistentFlags().StringVarP(&instrument, "inst", "i", "G", "Set the instrument type, G for guitar, B for Bass, U for Ukulele")
}

var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Print scale",
	Long:  `Print scale that you pass in argument. Example: C-D-E-F-G-A-B. Takes the first note as scale's root`,

	Run: func(cmd *cobra.Command, args []string) {

		scaleInput := args[0]

		scale, err := music.ParseScale(scaleInput, name)

		if err != nil {
			panic("Invalid Scale Input")
		}

		neck := neck.New(instrument)

		scaleAsString := render.RenderScale(neck, scale, from, to, mode)

		fmt.Print(scaleAsString)
	},
}
