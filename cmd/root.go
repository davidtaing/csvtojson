package cmd

import (
	"fmt"
	"io"
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
		r             io.Reader
		err           error
		csvFile       *os.File
		inputFromFile bool = input != ""
	)

	if inputFromFile {
		csvFile, err = app.OpenCSVFile(input)
		defer csvFile.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read input file: %s\n", input)
			return
		}

		r = csvFile
	} else {
		fmt.Fprintln(os.Stderr, "Reading from stdin")

		fi, _ := os.Stdin.Stat()

		if fi.Size() == 0 {
			fmt.Fprintln(os.Stderr, "No input detected from stdin")
			return
		}

		r = os.Stdin
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
