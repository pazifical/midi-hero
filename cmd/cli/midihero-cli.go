package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"github.com/pazifical/midi-hero/internal/midi"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "midi", "", "path to a midi file")
	flag.Parse()

	if filePath == "" {
		os.Exit(1)
	}

	fmt.Println(filePath)

	chart, err := midi.ImportFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = clonehero.WriteToFile(chart, fmt.Sprintf("%s.chart", filePath))
	if err != nil {
		log.Fatal(err)
	}
}
