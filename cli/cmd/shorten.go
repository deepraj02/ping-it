/*
Copyright © 2024 Deepraj Baidya deeprajbaidya@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shortenCmd represents the shorten command
var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Takes a String as an URL and returns shorten version of it.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shorten called")
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
