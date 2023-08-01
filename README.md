# CSV to JSON Converter

Converts CSV to JSON and outputs to stdout.

Pipe stdout to the `tee` CLI for better results.

## Tech

- Golang
- Cobra CLI

## Usage

There is an example data.csv file that you can use to run this project.

### Reading from CSV file

```
$ go run main.go -i data.csv

// combine with tee to output to file
$ go run main.go -i data.csv | tee output.json
```

### Reading from stdin

```
$ cat data.csv | go run main.go

// can also be outputted to file via tee
$ cat data.csv | go run main.go | tee output.json

```

## Things I Learnt

- All stdout writes are piped as output. So all log messages have to be written to `os.Stderr` instead. This is what the std log package uses under the hood.

- Marshalling data of unknown type to JSON

- io.Pipe is for piping variables across different goroutines, and are not releated to Unix Pipes
