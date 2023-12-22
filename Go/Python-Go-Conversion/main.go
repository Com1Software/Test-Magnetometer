package main

import (
	"fmt"
	"math"
	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

const (
	ADDRESS   = 0x1E
	CONFIG_A  = 0x00
	CONFIG_B  = 0x01
	MODE      = 0x02
	X_MSB     = 0x03
	Z_MSB     = 0x05
	Y_MSB     = 0x07
	declAngle = 0.22
)

func setup() {
	bus := embd.NewI2CBus(1)
	bus.WriteByteToReg(ADDRESS, CONFIG_A, 0x70)
	bus.WriteByteToReg(ADDRESS, CONFIG_B, 0x20)
	bus.WriteByteToReg(ADDRESS, MODE, 0x00)
}

func readRawData(addr byte) int {
	bus := embd.NewI2CBus(1)
	high, _ := bus.ReadByteFromReg(ADDRESS, addr)
	low, _ := bus.ReadByteFromReg(ADDRESS, addr+1)

	value := int(high<<8) + int(low)
	if value > 32768 {
		value = value - 65536
	}
	return value
}

func computeHeading(x, y int) float64 {
	headingRad := math.Atan2(float64(y), float64(x))

	headingRad += declAngle

	if headingRad < 0 {
		headingRad += 2 * math.Pi
	}

	if headingRad > 2*math.Pi {
		headingRad -= 2 * math.Pi
	}

	headingDeg := headingRad * (180.0 / math.Pi)
	return headingDeg
}

func main() {
	setup()
	for {
		x := readRawData(X_MSB)
		y := readRawData(Y_MSB)
		z := readRawData(Z_MSB)

		heading := computeHeading(x, y)
		fmt.Printf("heading %s %d\n", heading, z)

		time.Sleep(100 * time.Millisecond)
	}
}
