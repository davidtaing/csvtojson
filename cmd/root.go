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
		err error
		r   *os.File
	)

	if input != "" {
		r, err = app.OpenCSVFile(input)
		defer r.Close()
	} else {
		r = os.Stdin
		// todo check stdin input size
	}

	if err != nil {
		m := fmt.Sprintf("Failed to read from file: %s", input)
		fmt.Fprintln(os.Stderr, m)
	}

	err = app.ConvertCSVToJSON(r, os.Stdout)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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
