/*
Copyright © 2024 Deepraj Baidya deeprajbaidya@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pingit",
	Short: "PingIt is an URL Shortener CLI tool.",

	Long: `
	PingIt is a URL shortener service that provides both a command-line interface (CLI) and a web API, written in Go.
	This project is designed to offer a simple and efficient way to shorten URLs directly from your terminal or through HTTP requests.
	
	Created with ❤️ by @Deepraj02
	
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}