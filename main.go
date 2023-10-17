package main

import "flag"

func main() {
	var inputFile, outputFile string
	flag.StringVar(&inputFile, "inputFile", "", "Input file to read from")
	flag.StringVar(&outputFile, "outputFile", "", "Output file to write to")
	flag.Parse()

}
