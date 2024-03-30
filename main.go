package main

import (
	"log"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"github.com/pazifical/midi-hero/internal/midi"
)

func main() {
	chart, err := midi.ImportFile("testdata/midi.mid")
	if err != nil {
		log.Fatal(err)
	}

	err = clonehero.WriteToFile(chart, "testdata/midi.chart")
	if err != nil {
		log.Fatal(err)
	}
}
