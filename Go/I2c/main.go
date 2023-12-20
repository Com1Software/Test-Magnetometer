package main

import (
	"fmt"

	"golang.org/x/exp/io/i2c"
)

const (
	devaddr = 0x1E
	rega = 0x00
	regb = 0x01
	mode = 0x02
	x_axis = 0x03
	y_axis = 0x07
	z_axis = 0x05
	declination = -0.00669
	pi = 3.14159265359
)

func main() {
	x:=[]byte{}
	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x1e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Open %s\n", d)
	d.Read(x)
	fmt.Println(x)
}
