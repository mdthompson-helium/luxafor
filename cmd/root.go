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
		flagVal, err := cmd.Flags().GetBool("toggle");
		if err != nil {
			return
		}
		if flagVal {
			fmt.Println("Toggle DND");
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Toggle luxafor dnd status")
}


