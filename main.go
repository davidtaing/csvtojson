package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fd, err := os.Open("data.csv")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened the CSV file")
	defer fd.Close()

	r := csv.NewReader(fd)
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
