// CLI tool for compressing files with snappy.
//
// Please report any feature requests or bugs to https://github.com/suicidejack/snappycli/issues
package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/golang/snappy/snappy"
)

var (
	inputFilename  string
	outputFilename string
)

func init() {
	flag.StringVar(&inputFilename, "input-file", "", "[REQUIRED] the file to be compressed")
	flag.StringVar(&outputFilename, "output-file", "", "name of the file with the compressed contents - defaults to [input-file].snap")
}

func main() {
	validateFlags()

	var err error
	var in, out *os.File

	if in, err = os.Open(inputFilename); err != nil {
		log.Fatalf("unable to open input file %s: %s", inputFilename, err.Error())
	}
	bufin := bufio.NewReader(in)

	if out, err = os.Create(outputFilename); err != nil {
		log.Fatalf("unable to create output file %s: %s", outputFilename, err.Error())
	}
	writer := snappy.NewWriter(out)

	bufin.WriteTo(writer)
	in.Close()
	out.Close()
}

func validateFlags() {
	flag.Parse()

	if inputFilename == "" {
		log.Fatal("you must supply an input file")
	}
	if outputFilename == "" {
		outputFilename = inputFilename + ".snap"
	}
}
