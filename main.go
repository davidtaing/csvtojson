package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("data.csv")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened the CSV file")
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

func pipe() {
	_, w := io.Pipe()

	go func() {
		_, err := fmt.Fprint(w, "Hello World\n")

		if err != nil {
			fmt.Println(err)
		}
		w.Close()
	}()
}
