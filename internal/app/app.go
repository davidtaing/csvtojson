package app

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func ConvertCSVToJSON(path string) error {
	ext := filepath.Ext(path)

	if ext != ".csv" {
		var m = fmt.Sprintf("File %s is not a .csv file\n", path)
		fmt.Printf(m)
		return errors.New(m)
	}

	records, err := readCSV(path)
	if err != nil {
		fmt.Printf("Failed to read csv %s\n", path)
		return err
	}

	jsonBytes, err := marshalToJSON(records)
	if err != nil {
		fmt.Printf("Failed to marshal csv data")
		return err
	}

	_, err = os.Stdout.Write(jsonBytes)

	if err != nil {
		fmt.Println("Error writing data to json", err)
		return err
	}

	return nil
}

func readCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Fprintln(os.Stderr, "Successfully opened the CSV file")

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return records, nil
}

func marshalToJSON(records [][]string) ([]byte, error) {
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
		fmt.Println("Error marshaling data:", err)
		return nil, err
	}

	return jsonBytes, nil
}
