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
	Run:   CSVToJSONCommand,
}

func CSVToJSONCommand(cmd *cobra.Command, args []string) {
	var (
		records [][]string
		err     error
	)

	if input == "" {
		fmt.Fprintln(os.Stderr, "reading csv from stdin")
		records, err = app.ReadCSVFromStdin()
	} else {
		fmt.Fprintln(os.Stderr, "reading csv from file")
		records, err = app.ReadCSVFromFile(input)
	}

	if err != nil {
		m := fmt.Sprintf("Read failed %s", err)
		fmt.Fprintln(os.Stderr, m, err)
		return
	}

	err = app.ConvertCSVToJSON(records)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to marshal csv data")
		return
	}

	// Bump terminal prompt to sit on new line
	fmt.Print("\n")
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
