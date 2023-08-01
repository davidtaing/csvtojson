package cmd

import (
	"fmt"
	"os"

	"github.com/davidtaing/csvtojson/internal/app"
	"github.com/spf13/cobra"
)

var (
	input string
)

var rootCmd = &cobra.Command{
	Use:   "csvtojson",
	Short: "Convert a CSV file to JSON. Outputted to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		source := "stdin"
		records, err := app.ReadCSVFromStdin()

		if len(records) == 0 {
			source = fmt.Sprintf("file: %s", input)
			records, err = app.ReadCSVFromFile(input)
		}

		if err != nil {
			m := fmt.Sprintf("Error reading from %s", source)
			fmt.Fprintln(os.Stderr, m)
			return
		}

		err = app.ConvertCSVToJSON(records)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to marshal csv data")
			return
		}

		// Bump terminal prompt to sit on new line
		fmt.Print("\n")
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
