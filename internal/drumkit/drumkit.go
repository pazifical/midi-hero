package drumkit

import "slices"

type Part string

const Kick Part = "Kick"
const Snare Part = "Snare"
const SnareRimshot Part = "Snare Rimshot"
const Tom1 Part = "Tom 1"
const Tom2 Part = "Tom 2"
const Tom3 Part = "Tom 3"
const HiHatClosed Part = "HiHat Closed"
const HiHatOpen Part = "HiHat Open"
const Crash1Left Part = "Crash 1 Left"
const China1Left Part = "China 1 Left"
const SplashCenter Part = "Splash Center"
const China2CenterRight Part = "China 2 Center Right"
const China2Right Part = "China 2 Right"
const Crash2Right Part = "Crash 2 Right"
const Crash3Right Part = "Crash 3 Right"
const China3Right Part = "China 3 Right"
const Ride Part = "Ride"
const RideBell Part = "Ride Bell"

var cymbals = []Part{
	HiHatClosed,
	HiHatOpen,
	Crash1Left,
	China1Left,
	SplashCenter,
	China2CenterRight,
	China2Right,
	Crash2Right,
	Crash3Right,
	China3Right,
	Ride,
	RideBell,
}

func IsCymbal(part Part) bool {
	return slices.Contains(cymbals, part)
}
