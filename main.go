package main

import (
	"flag"

	"github.com/merschformann/rendercube/rendering"
)

func main() {
	var inputFile, outputFile string
	flag.StringVar(&inputFile, "input", "", "Input file to read from")
	flag.StringVar(&outputFile, "output", "", "Output file to write to")
	flag.Parse()

	input, err := rendering.ReadInput(inputFile)
	if err != nil {
		panic(err)
	}

	err = rendering.Convert(input, outputFile)
	if err != nil {
		panic(err)
	}
}
