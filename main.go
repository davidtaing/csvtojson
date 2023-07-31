package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	f := readCSV("data.csv")
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

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
		return
	}

	_, err = os.Stdout.Write(jsonBytes)

	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}

func readCSV(path string) *os.File {
	f, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Successfully opened the CSV file")

	return f
}
