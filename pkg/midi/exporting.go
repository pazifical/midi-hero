package midi

import (
	"fmt"
	"math"

	"github.com/pazifical/midi-hero/pkg/clonehero"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func ExportToFile(chart clonehero.Chart, filepath string) error {
	var tr smf.Track
	var clock = smf.MetricTicks(chart.Song.Resolution)

	ts := chart.SyncTrack.TimeSignatures[0]
	first := uint8(ts.Numerator)
	second := uint8(math.Pow(float64(ts.Denominator), 2))
	tr.Add(0, smf.MetaMeter(first, second))

	bpm := chart.SyncTrack.Tempos[0]
	tr.Add(0, smf.MetaTempo(float64(bpm.MilliBPM/1000)))

	tr.Add(0, smf.MetaInstrument("Drums"))

	notePositions := make(map[int][]clonehero.Note, 0)
	for _, noteEvent := range chart.ExpertDrums.Values {
		_, ok := notePositions[noteEvent.Position]
		if !ok {
			notePositions[noteEvent.Position] = make([]clonehero.Note, 0)
		}
		notePositions[noteEvent.Position] = append(notePositions[noteEvent.Position], noteEvent.Note)
	}
	fmt.Println(notePositions)

	previousPosition := 0
	for position, notes := range notePositions {
		for i, note := range notes {
			fmt.Println(note)
			// TODO
			if i > 0 {
				tr.Add(0, midi.NoteOn(0, midi.C(3), 120))
			} else {
				tr.Add(uint32(position-previousPosition), midi.NoteOn(0, midi.C(3), 120))
			}
		}
	}

	tr.Add(0, midi.NoteOn(0, midi.Ab(3), 120))
	tr.Add(clock.Ticks8th(), midi.NoteOn(0, midi.C(4), 120))
	tr.Add(clock.Ticks4th()*2, midi.NoteOff(0, midi.Ab(3)))
	tr.Add(0, midi.NoteOff(0, midi.C(4)))
	tr.Close(0)

	// create the SMF and add the tracks
	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)

	err := s.WriteFile(filepath)
	if err != nil {
		return err
	}

	return nil
}
