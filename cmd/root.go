package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

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
		defer csvFile.Close()

		if err != nil {
			return err
		}

		r = csvFile
	} else {
		fmt.Fprintln(os.Stderr, "Reading from stdin")

		// Use a goroutine with a timeout to read from stdin
		done := make(chan bool)
		var buf bytes.Buffer
		go func() {
			_, err := io.Copy(&buf, os.Stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read from stdin: %v\n", err)
				done <- false
			} else {
				done <- true
			}
		}()

		select {
		case success := <-done:
			if !success {
				return errors.New("Failed to read from stdin")
			}
			r = &buf
		case <-time.After(1 * time.Second):
			return errors.New("Failed to read from stdin")
		}
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
