/*
Copyright © 2024 Deepraj Baidya deeprajbaidya@gmail.com
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/deepraj02/pingit-cli/helpers"
	"github.com/spf13/cobra"
)

// shortenCmd represents the shorten command
var shortenCmd = &cobra.Command{
	Use:   "shorten [url]",
	Short: "Takes a String as an URL and returns shorten version of it.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		if !govalidator.IsURL(url) {
			fmt.Println("Invalid URL provided.")
			os.Exit(1)
		}

		// Prepare the request body

		requestBody, err := json.Marshal(map[string]string{
			"url": url,
		})
		if err != nil {
			fmt.Println("Error preparing request:", err)
			os.Exit(1)
		}

		// Define the API endpoint
		/// TODO: Remove HardCoded Data
		const apiURL = "http://localhost:3000/api/v1/"

		// Create a new request
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("Error creating request:", err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			os.Exit(1)
		}
		defer response.Body.Close()

		//If Status Okay >?
		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: received status code %d\n", response.StatusCode)
			os.Exit(1)
		}

		// Decode and Read the response
		var apiResponse helpers.APIResponse
		if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
			fmt.Println("Error decoding response:", err)
			os.Exit(1)
		}

		fmt.Println("URL shortened successfully.")
		fmt.Println("Shortened URL:", apiResponse.ShortURL)
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
}
