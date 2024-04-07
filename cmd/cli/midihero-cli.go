package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pazifical/midi-hero/pkg/clonehero"
	"github.com/pazifical/midi-hero/pkg/midi"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "midi", "", "path to a midi file")
	flag.Parse()

	if filePath == "" {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	chart, err := midi.ImportFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = clonehero.WriteToFile(chart, fmt.Sprintf("%s.chart", filePath))
	if err != nil {
		log.Fatal(err)
	}
}
