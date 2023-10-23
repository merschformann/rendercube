package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"

	"github.com/merschformann/rendercube/rendering"
)

func main() {
	var inputFile, outputFile string
	flag.StringVar(&inputFile, "input", "", "Input file to read from")
	flag.StringVar(&outputFile, "output", "", "Output file to write to")
	flag.Parse()

	input, err := readInput(inputFile)
	if err != nil {
		panic(err)
	}

	output, err := rendering.Convert(input)
	if err != nil {
		panic(err)
	}

	err = writeOutput(outputFile, output)
	if err != nil {
		panic(err)
	}
}

// readInput reads the input from the given file (empty string for stdin).
func readInput(file string) (rendering.Input, error) {
	var content []byte
	var err error
	if file != "" {
		content, err = os.ReadFile(file)
		if err != nil {
			return rendering.Input{}, err
		}
	} else {
		content, err = io.ReadAll(os.Stdin)
		if err != nil {
			return rendering.Input{}, err
		}
	}
	var input rendering.Input
	err = json.Unmarshal(content, &input)
	if err != nil {
		return rendering.Input{}, err
	}
	return input, nil
}

// writeOutput writes the output to the given file (empty string for stdout).
func writeOutput(file string, output string) error {
	if file != "" {
		err := os.WriteFile(file, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else {
		_, err := os.Stdout.WriteString(output)
		if err != nil {
			return err
		}
	}
	return nil
}
