package midi

import (
	"bytes"
	"fmt"
	"math"
	"os"

	"github.com/pazifical/midi-hero/internal/clonehero"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

func ImportFile(filePath string) (clonehero.Chart, error) {
	defer midi.CloseDriver()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return clonehero.Chart{}, nil
	}

	reader := bytes.NewReader(data)

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

func run() {
	defer midi.CloseDriver()

	// create a SMF
	rd := bytes.NewReader(mkSMF())

	// read and play it
	smf.ReadTracksFrom(rd).Do(func(ev smf.TrackEvent) {
		fmt.Printf("track %v @%vms %s\n", ev.TrackNo, ev.AbsMicroSeconds/1000, ev.Message)
	})
}

// makes a SMF and returns the bytes
func mkSMF() []byte {
	var (
		bf    bytes.Buffer
		clock = smf.MetricTicks(96) // resolution: 96 ticks per quarternote 960 is also common
		tr    smf.Track
	)

	// first track must have tempo and meter informations
	tr.Add(0, smf.MetaMeter(3, 4))
	tr.Add(0, smf.MetaTempo(140))
	tr.Add(0, smf.MetaInstrument("Brass"))
	tr.Add(0, midi.ProgramChange(0, gm.Instr_BrassSection.Value()))
	tr.Add(0, midi.NoteOn(0, midi.Ab(3), 120))
	tr.Add(clock.Ticks8th(), midi.NoteOn(0, midi.C(4), 120))
	// duration: a quarter note (96 ticks in our case)
	tr.Add(clock.Ticks4th()*2, midi.NoteOff(0, midi.Ab(3)))
	tr.Add(0, midi.NoteOff(0, midi.C(4)))
	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()
}
