/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "luxafor",
	Short: "Luxafor CLI tool",
	Long: `A tool to turn on DND via a cli tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		on, err := cmd.Flags().GetBool("dnd-on");
		if err != nil {
			return
		}
		luxafor := NewLuxafor()
		defer luxafor.Close()

		if on {
			fmt.Println("DND on");
			luxafor.Colour(LedAll, 8, 0 , 0, 0)
			return
		}

		if on {
			fmt.Println("DND off");
			luxafor.Colour(LedAll, 0, 8 , 0, 0)
			return
		}
		fmt.Println("Hello World");
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("dnd-on", "o", false, "Turn dnd on")
	rootCmd.Flags().BoolP("dnd-off", "f", false, "Turn dnd off")
}


