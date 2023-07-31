# CSV to JSON Converter

Converts CSV to JSON and outputs to stdout.

Pipe stdout to the `tee` CLI for better results.

## Tech

- Golang
- Cobra CLI

## Usage

There is an example data.csv file that you can use to run this project.

To run this CLI, clone the project and run the following command.

```
$ go run main.go -i data.csv

// combine with unix pipe & tee <output_filepath>
$ go run main.go -i data.csv | tee output.json
```

## Things I Learnt

- All stdout writes are piped as output. So all log messages have to be written to `os.Stderr` instead. This is what the std log package uses under the hood.

- Marshalling data of unknown type to JSON
