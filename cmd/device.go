package cmd

import (
	"log"

	"github.com/karalabe/hid"
)

const vendorId = uint16(0x4d8)   // Microchip Technology Inc.
const productId = uint16(0xf372) // LUXAFOR FLAG

type Luxafor struct {
	Device *hid.Device
}

func newDevice() *hid.Device {

	devInfo := hid.Enumerate(vendorId, productId)

	if len(devInfo) < 1 {
		log.Fatalf("no devices found matching VID %v and PID %v", vendorId, productId)
	}
	if len(devInfo) > 1 {
		log.Fatalf("More than one device found matching VID %v and PID %v", vendorId, productId)
	}

	dev, err := devInfo[0].Open()
	if err != nil {
		log.Fatalf("Failed to open HID device, err: %s", err)
	}

	return dev
}