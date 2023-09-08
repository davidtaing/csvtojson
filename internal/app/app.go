package app

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func OpenCSVFile(p string) (*os.File, error) {
	fmt.Fprintf(os.Stderr, "Reading from CSV file: %s\n", p)

	f, err := os.Open(p)
	return f, err
}

func ConvertCSVToJSON(r io.Reader, w io.Writer) error {
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read from input.")
		return err
	}

	jsonBytes, err := MarshalToJSON(records)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to marshal csv data")
		return err
	}

	_, err = w.Write(jsonBytes)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing data to json", err)
		return err
	}

	return nil
}

func ReadCSVFromStdin() (io.Reader, error) {
	fmt.Fprintln(os.Stderr, "Reading from stdin")

	done := make(chan bool)
	var buf bytes.Buffer
	go func() {
		_, err := io.Copy(&buf, os.Stdin)
		if err != nil {
			done <- false
		} else {
			done <- true
		}
	}()

	select {
	case success := <-done:
		if !success {
			return nil, errors.New("Failed to read from stdin")
		}

		return &buf, nil
	case <-time.After(1 * time.Second):
		return nil, errors.New("Failed to read from stdin")
	}
}

func MarshalToJSON(records [][]string) ([]byte, error) {
	var data []map[string]string

	if len(records) > 0 {
		columnNames := records[0]

		for _, row := range records[1:] {
			record := make(map[string]string)
			for i, value := range row {
				if i < len(columnNames) {
					record[columnNames[i]] = value
				}
			}
			data = append(data, record)
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling data:", err)
		return nil, err
	}

	return jsonBytes, nil
}
