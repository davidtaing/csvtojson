package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var p = "data.csv"

	records, err := readCSV(p)
	if err != nil {
		fmt.Printf("Failed to read csv %s\n", p)
		return
	}

	jsonBytes, err := marshalToJSON(records)
	if err != nil {
		fmt.Printf("Failed to marshal csv data")
		return
	}

	_, err = os.Stdout.Write(jsonBytes)

	if err != nil {
		fmt.Println("Error writing data to json", err)
		return
	}
}

func readCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully opened the CSV file")

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
