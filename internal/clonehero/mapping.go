package clonehero

import "github.com/pazifical/midi-hero/internal/drumkit"

type Note uint8

const Kick Note = 0
const Red Note = 1
const Yellow Note = 2
const Blue Note = 3
const Orange Note = 4
const Green Note = 5
const RedAccent Note = 34
const YellowAccent Note = 35
const BlueAccent Note = 36
const OrangeAccent Note = 37
const GreenAccent Note = 38
const RedGhost Note = 40
const YellowGhost Note = 41
const BlueGhost Note = 42
const OrangeGhost Note = 43
const GreenGhost Note = 44
const YellowCymbal Note = 66
const BlueCymbal Note = 67
const GreenCymbal Note = 68

func StylesForPart(part drumkit.Part) []Note {
	notes, ok := styleMapping[part]
	if !ok {
		return make([]Note, 0)
	}
	return notes
}

var styleMapping = map[drumkit.Part][]Note{
	drumkit.Kick:              {Kick},
	drumkit.Snare:             {Red},
	drumkit.SnareRimshot:      {Red, RedAccent},
	drumkit.HiHatOpen:         {Yellow, YellowCymbal},
	drumkit.HiHatClosed:       {Yellow, YellowCymbal},
	drumkit.Crash1Left:        {Yellow, YellowCymbal},
	drumkit.China1Left:        {Yellow, YellowAccent, YellowCymbal},
	drumkit.SplashCenter:      {Blue, BlueGhost, BlueCymbal},
	drumkit.China2CenterRight: {Blue, BlueAccent, BlueCymbal},
	drumkit.Crash2Right:       {Blue, BlueCymbal},
	drumkit.Ride:              {Blue, BlueCymbal},
	drumkit.RideBell:          {Blue, BlueAccent, BlueCymbal},
	drumkit.Crash3Right:       {Green, GreenCymbal},
	drumkit.China3Right:       {Green, GreenAccent, GreenCymbal},
}
