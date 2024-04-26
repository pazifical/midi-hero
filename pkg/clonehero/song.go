package clonehero

import (
	"fmt"
	"strings"
)

type Chart struct {
	Song        SongSection
	SyncTrack   SyncTrackSection
	Events      EventSection
	ExpertDrums ExpertDrumSection
}

func (ch *Chart) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("[Song]\n%s", ch.Song.String()))
	builder.WriteString(fmt.Sprintf("[SyncTrack]\n%s", ch.SyncTrack.String()))
	builder.WriteString("[Events]\n{\n}\n") // TODO: implement
	builder.WriteString(fmt.Sprintf("[ExpertDrums]\n%s", ch.ExpertDrums.String()))
	return builder.String()
}

type SongSection struct {
	Offset       int
	Resolution   int
	Player2      string
	Difficulty   int
	PreviewStart int
	PreviewEnd   int
	Genre        string
	MediaType    string
}

func (ss *SongSection) String() string {
	var builder strings.Builder
	builder.WriteString("{\n")
	builder.WriteString(fmt.Sprintf("  Offset = %d\n", ss.Offset))
	builder.WriteString(fmt.Sprintf("  Resolution = %d\n", ss.Resolution))
	builder.WriteString(fmt.Sprintf("  Player2 = %s\n", ss.Player2))
	builder.WriteString(fmt.Sprintf("  Difficulty = %d\n", ss.Difficulty))
	builder.WriteString(fmt.Sprintf("  PreviewStart = %d\n", ss.PreviewStart))
	builder.WriteString(fmt.Sprintf("  PreviewEnd = %d\n", ss.PreviewEnd))
	builder.WriteString(fmt.Sprintf("  Genre = %s\n", ss.Genre))
	builder.WriteString(fmt.Sprintf("  MediaType = %s\n", ss.MediaType))
	builder.WriteString("}\n")
	return builder.String()
}

type SyncTrackSection struct {
	TimeSignatures []TimeSignature
	Tempos         []Tempo
}

type EventSection struct {
}

func (sts *SyncTrackSection) String() string {
	var builder strings.Builder
	builder.WriteString("{\n")
	for _, ts := range sts.TimeSignatures {
		builder.WriteString(fmt.Sprintf("  %d = TS %d %d\n", ts.Position, ts.Numerator, ts.Denominator))
	}
	for _, ts := range sts.Tempos {
		builder.WriteString(fmt.Sprintf("  %d = B %d\n", ts.Position, ts.MilliBPM))
	}
	builder.WriteString("}\n")
	return builder.String()
}

// <Position> = B <Tempo>
type Tempo struct {
	Position int
	MilliBPM int
}

// <Position> = TS <Numerator> <Denominator exponent>
// Numerator is the numerator to use for the time signature.
// Denominator exponent is optional, and specifies a power of 2 to use for the denominator of the time signature.
// Example:
// 0 = TS 4       // 4/4
// 0 = TS 4 2     // Also 4/4
// 768 = TS 7 4   // 7/16
// 1104 = TS 3 3  // 3/8
type TimeSignature struct {
	Position    int
	Numerator   int
	Denominator int
}

func NewTimeSignature(position, numerator, denominator int) TimeSignature {
	if denominator <= 0 {
		denominator = 2
	}
	return TimeSignature{
		Position:    position,
		Numerator:   numerator,
		Denominator: denominator,
	}
}

type ExpertDrumSection struct {
	Values []NoteEvent
}

func (eds *ExpertDrumSection) String() string {
	var builder strings.Builder
	builder.WriteString("{\n")
	for _, value := range eds.Values {
		builder.WriteString(fmt.Sprintf("  %d = N %d %d\n", value.Position, value.Note, value.Length))
	}
	builder.WriteString("}\n")
	return builder.String()
}

// <Position> = N <Type> <Length>
// Type is the type number of this note/modifier.
// Length is the length of this note in ticks.
// This value typically doesn't do anything for modifiers.
type NoteEvent struct {
	Position int
	Note     Note
	Length   int
}

func (de *NoteEvent) String() string {
	return fmt.Sprintf("%d = N %d %d", de.Position, de.Note, de.Length)
}
