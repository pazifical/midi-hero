package midi

import "github.com/pazifical/midi-hero/internal/drumkit"

type MidiNote uint8

var DrumMapping = map[MidiNote]drumkit.Part{
	35:  drumkit.Kick,
	36:  drumkit.Kick,
	37:  drumkit.Snare,
	38:  drumkit.Snare,
	40:  drumkit.Snare,
	41:  drumkit.Tom1,
	42:  drumkit.HiHatClosed,
	43:  drumkit.Tom3,
	44:  drumkit.HiHatOpen,
	45:  drumkit.Tom3,
	46:  drumkit.HiHatOpen,
	47:  drumkit.Tom2,
	48:  drumkit.Tom1,
	49:  drumkit.Crash3Right,
	50:  drumkit.Tom3,
	51:  drumkit.Ride,
	52:  drumkit.China3Right,
	53:  drumkit.RideBell,
	55:  drumkit.SplashCenter,
	56:  drumkit.RideBell,
	57:  drumkit.Crash2Right,
	59:  drumkit.Ride,
	91:  drumkit.SnareRimshot,
	92:  drumkit.HiHatOpen,
	93:  drumkit.Ride,
	99:  drumkit.RideBell,
	102: drumkit.RideBell,
}
