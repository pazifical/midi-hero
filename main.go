package main

import (
	"log"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"github.com/pazifical/midi-hero/internal/midi"
)

func main() {
	chart, err := midi.ImportFile("testdata/song2.mid")
	if err != nil {
		log.Fatal(err)
	}

	err = clonehero.WriteToFile(chart, "testdata/song2.chart")
	if err != nil {
		log.Fatal(err)
	}
}
