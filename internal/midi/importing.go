package midi

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func ImportFromReader(reader io.Reader) (clonehero.Chart, error) {
	noteEvents := make([]clonehero.NoteEvent, 0)
	tempos := make([]clonehero.Tempo, 0)
	timeSignatures := make([]clonehero.TimeSignature, 0)

	smf.ReadTracksFrom(reader).Do(func(ev smf.TrackEvent) {
		fmt.Printf("track %v %d @%vms %s\n", ev.TrackNo, ev.AbsTicks, ev.AbsMicroSeconds/1000, ev.Message)

		var channel uint8
		var key uint8
		var velocity uint8
		isNoteOn := ev.Message.GetNoteOn(&channel, &key, &velocity)
		if isNoteOn {
			part, ok := DrumMapping[MidiNote(key)]
			if !ok {
				return
			}

			notes := clonehero.StylesForPart(part)
			for _, note := range notes {
				noteEvents = append(noteEvents, clonehero.NoteEvent{
					Position: int(ev.AbsTicks),
					Note:     note,
				})
			}
			return
		}

		var bpm float64
		isTempo := ev.Message.GetMetaTempo(&bpm)
		if isTempo {
			tempos = append(tempos, clonehero.Tempo{
				Position: int(ev.AbsTicks),
				MilliBPM: int(bpm * 1000),
			})
			return
		}

		var numerator uint8
		var denominator uint8
		var clocksPerClick uint8
		var demiSemiQuaverPerQuarter uint8
		isTimeSig := ev.Message.GetMetaTimeSig(&numerator, &denominator, &clocksPerClick, &demiSemiQuaverPerQuarter)
		if isTimeSig {
			denominatorExp := math.Sqrt(float64(denominator))
			timeSignatures = append(timeSignatures, clonehero.NewTimeSignature(
				int(ev.AbsTicks),
				int(numerator),
				int(denominatorExp),
			))
			return
		}

	})

	chart := clonehero.Chart{
		ExpertDrums: clonehero.ExpertDrumSection{
			Values: noteEvents,
		},
		Song: clonehero.SongSection{
			Offset:       0,
			Resolution:   192,
			Player2:      "Player2",
			Difficulty:   0,
			PreviewStart: 0,
			PreviewEnd:   0,
			Genre:        "Genre",
			MediaType:    "MediaType",
		},
		Events: clonehero.EventSection{},
		SyncTrack: clonehero.SyncTrackSection{
			TimeSignatures: timeSignatures,
			Tempos:         tempos,
		},
	}

	return chart, nil
}

func ImportFile(filePath string) (clonehero.Chart, error) {
	defer midi.CloseDriver()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return clonehero.Chart{}, nil
	}

	reader := bytes.NewReader(data)

	return ImportFromReader(reader)

	// noteEvents := make([]clonehero.NoteEvent, 0)
	// tempos := make([]clonehero.Tempo, 0)
	// timeSignatures := make([]clonehero.TimeSignature, 0)

	// smf.ReadTracksFrom(reader).Do(func(ev smf.TrackEvent) {
	// 	fmt.Printf("track %v %d @%vms %s\n", ev.TrackNo, ev.AbsTicks, ev.AbsMicroSeconds/1000, ev.Message)

	// 	var channel uint8
	// 	var key uint8
	// 	var velocity uint8
	// 	isNoteOn := ev.Message.GetNoteOn(&channel, &key, &velocity)
	// 	if isNoteOn {
	// 		part, ok := DrumMapping[MidiNote(key)]
	// 		if !ok {
	// 			return
	// 		}

	// 		notes := clonehero.StylesForPart(part)
	// 		for _, note := range notes {
	// 			noteEvents = append(noteEvents, clonehero.NoteEvent{
	// 				Position: int(ev.AbsTicks),
	// 				Note:     note,
	// 			})
	// 		}
	// 		return
	// 	}

	// 	var bpm float64
	// 	isTempo := ev.Message.GetMetaTempo(&bpm)
	// 	if isTempo {
	// 		tempos = append(tempos, clonehero.Tempo{
	// 			Position: int(ev.AbsTicks),
	// 			MilliBPM: int(bpm * 1000),
	// 		})
	// 		return
	// 	}

	// 	var numerator uint8
	// 	var denominator uint8
	// 	var clocksPerClick uint8
	// 	var demiSemiQuaverPerQuarter uint8
	// 	isTimeSig := ev.Message.GetMetaTimeSig(&numerator, &denominator, &clocksPerClick, &demiSemiQuaverPerQuarter)
	// 	if isTimeSig {
	// 		denominatorExp := math.Sqrt(float64(denominator))
	// 		timeSignatures = append(timeSignatures, clonehero.NewTimeSignature(
	// 			int(ev.AbsTicks),
	// 			int(numerator),
	// 			int(denominatorExp),
	// 		))
	// 		return
	// 	}

	// })

	// chart := clonehero.Chart{
	// 	ExpertDrums: clonehero.ExpertDrumSection{
	// 		Values: noteEvents,
	// 	},
	// 	Song: clonehero.SongSection{
	// 		Offset:       0,
	// 		Resolution:   192,
	// 		Player2:      "Player2",
	// 		Difficulty:   0,
	// 		PreviewStart: 0,
	// 		PreviewEnd:   0,
	// 		Genre:        "Genre",
	// 		MediaType:    "MediaType",
	// 	},
	// 	Events: clonehero.EventSection{},
	// 	SyncTrack: clonehero.SyncTrackSection{
	// 		TimeSignatures: timeSignatures,
	// 		Tempos:         tempos,
	// 	},
	// }

	// return chart, nil
}
