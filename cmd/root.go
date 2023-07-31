package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/davidtaing/csvtojson/internal/app"
)

var (
	input string
)

var rootCmd = &cobra.Command{
	Use:   "csvtojson",
	Short: "Convert a CSV file to JSON. Outputted to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		err := app.ConvertCSVToJSON(input)

		if err != nil {
			return
		}

		// add newline to JSON output
		fmt.Println("")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&input, "input", "i", "", "input CSV file path")
}
