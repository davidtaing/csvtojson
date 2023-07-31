package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f := readCSV("data.csv")
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, row := range records {
		fmt.Println(row)
	}
}

func readCSV(path string) *os.File {
	f, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened the CSV file")

	return f
}
