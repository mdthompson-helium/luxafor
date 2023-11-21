package cmd

import (
	"log"
	"math/rand"
)

type Led byte

const (
	LedAll Led = 0xff
	LedA   Led = 0x41
	LedB   Led = 0x42
	Led1   Led = 0x01
	Led2   Led = 0x02
	Led3   Led = 0x03
	Led4   Led = 0x04
	Led5   Led = 0x05
	Led6   Led = 0x06
)

// Wave is one of 5 different wave patterns
type Wave byte

// The Wave patterns are fairly similar but with slightly different
// behaviours
const (
	Wave1 Wave = 0x01
	Wave2 Wave = 0x02
	Wave3 Wave = 0x03
	Wave4 Wave = 0x04
	Wave5 Wave = 0x05
)

// Pattern specifies a built-in pattern that affects all LEDs.
type Pattern byte

/* Patterns come in 8 flavours:

Pattern1 - Traffic lights
Pattern2 - Colour walk
Pattern3 - Random pattern
Pattern4 - Random fading pattern
Pattern5 - Police pattern
Pattern6 - Random quick fade pattern
Pattern7 - Colourful police pattern
Pattern8 - Quick rainbow pattern
*/
const (
	Pattern1 Pattern = 0x01
	Pattern2 Pattern = 0x02
	Pattern3 Pattern = 0x03
	Pattern4 Pattern = 0x04
	Pattern5 Pattern = 0x05
	Pattern6 Pattern = 0x06
	Pattern7 Pattern = 0x07
	Pattern8 Pattern = 0x08
)

// ProdCode is a "Productivity" mode is supposed to work but it's present in some of
// the documentation.
// It seems to be a "quick" way of setting individual colours based on letters like 'G'
// for green and 'R' for red. It has to be enabled before it will accept these codes.
type ProdCode byte

// Enable will enable this mode, Off seems to turn off all LEDs and
// Disable leaves productivity mode
const (
	Enable  ProdCode = 0x45
	Disable ProdCode = 0x44
	Red     ProdCode = 0x52
	Green   ProdCode = 0x47
	Blue    ProdCode = 0x42
	Cyan    ProdCode = 0x43
	Magenta ProdCode = 0x4d
	Yellow  ProdCode = 0x59
	White   ProdCode = 0x57
	Off     ProdCode = 0x4f
)

// RandomColour returns a random number between 0-255
func RandomColour() uint8 {
	return uint8(rand.Intn(0xff))
}

// NewLuxafor returns a new Luxafor device
func NewLuxafor() Luxafor {

	device := newDevice()

	return Luxafor{
		Device: device,
	}
}

// Close will close the connection to the USB device
func (l *Luxafor) Close() {
	l.Device.Close()
}

func (l *Luxafor) writeCommand(command []byte) error {
	_, err := l.Device.Write(command)
	if err != nil {
		log.Printf("Error writing data: %s", err)
	}
	return err
}

// Colour sets a RGB colour for specified LED(s)
func (l *Luxafor) Colour(led Led, red uint8, green uint8, blue uint8, fadeTime uint8) error {
	data := []byte{0x01, byte(led), red, green, blue, fadeTime, 0x0, 0x0}
	if fadeTime > 0 {
		data[1] = 0x02
	}
	return l.writeCommand(data)
}

// Strobe activates the strobe functionality for specified LED(s).
// Takes RGB values, speed and number of repetitions.
func (l *Luxafor) Strobe(led Led, red uint8, green uint8, blue uint8, speed uint8, repeat uint8) error {
	data := []byte{0x03, byte(led), red, green, blue, speed, 0x0, repeat}
	return l.writeCommand(data)
}

// Wave activates one of 5 wave patterns for the specified RBG colours.
// Speed and number of repetitions configurable. A repeat of 0 will repeat forever.
// Affects all LEDs.
func (l *Luxafor) Wave(wave Wave, red uint8, green uint8, blue uint8, speed uint8, repeat uint8) error {
	data := []byte{0x04, byte(wave), red, green, blue, 0x0, repeat, speed}
	return l.writeCommand(data)
}

// Pattern will activate one of 8 pre-programmed patterns.
// Repeat determines how many times it will run through the program,
// a repeat of 0 is forever.
// Affects all LEDs
func (l *Luxafor) Pattern(pattern Pattern, repeat uint8) error {
	data := []byte{0x06, byte(pattern), repeat, 0x0, 0x0, 0x0, 0x0, 0x0}
	return l.writeCommand(data)
}