package main

import (
	"fmt"

	"golang.org/x/exp/io/i2c"
)

func main() {
	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x39)
	if err != nil {
		panic(err)
	}
	fmt.Printf("first return %s", d)
	// opens a 10-bit address
	d, err = i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, i2c.TenBit(0x78))
	if err != nil {
		panic(err)
	}
	fmt.Printf("to bit return %s", d)

}
