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
	RunE:  CSVToJSONCommand,
}

func CSVToJSONCommand(cmd *cobra.Command, args []string) error {
	var (
		r             io.Reader
		err           error
		csvFile       *os.File
		inputFromFile bool = input != ""
	)

	if inputFromFile {
		csvFile, err = app.OpenCSVFile(input)
		r = csvFile
		defer csvFile.Close()
	} else {
		r, err = app.ReadCSVFromStdin()
	}

	if err != nil {
		return err
	}

	err = app.ConvertCSVToJSON(r, os.Stdout)

	if err != nil {
		return err
	}

	// Pad output so terminal return sits on new line
	fmt.Println()
	fmt.Fprintf(os.Stderr, "Exiting\n")
	return nil
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
