package midi

import "github.com/pazifical/midi-hero/internal/drumkit"

type MidiNote uint8

var DrumMapping = map[MidiNote]drumkit.Part{
	35:  drumkit.Kick,
	36:  drumkit.Kick,
	37:  drumkit.Snare,
	38:  drumkit.Snare,
	91:  drumkit.SnareRimshot,
	48:  drumkit.Tom1,
	47:  drumkit.Tom2,
	43:  drumkit.Tom3,
	45:  drumkit.Tom3,
	50:  drumkit.Tom3,
	42:  drumkit.HiHatClosed,
	44:  drumkit.HiHatOpen,
	46:  drumkit.HiHatOpen,
	92:  drumkit.HiHatOpen,
	55:  drumkit.SplashCenter,
	57:  drumkit.Crash2Right,
	49:  drumkit.Crash3Right,
	52:  drumkit.China3Right,
	51:  drumkit.Ride,
	93:  drumkit.Ride,
	53:  drumkit.RideBell,
	56:  drumkit.RideBell,
	99:  drumkit.RideBell,
	102: drumkit.RideBell,
}
