package main

import (
	"fmt"

	"github.com/d2r2/go-i2c"
)

const (
	devaddr     = 0x1E
	rega        = 0x00
	regb        = 0x01
	mode        = 0x02
	x_axis      = 0x03
	y_axis      = 0x07
	z_axis      = 0x05
	declination = -0.00669
	pi          = 3.14159265359
)

func main() {
	fmt.Println("Start")

	i2c, err := i2c.NewI2C(0x0d, 1)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(i2c)
	buf1 := make([]byte, 10)
	x, erra := i2c.ReadBytes(buf1)
	if erra != nil {
		fmt.Printf("Error %s\n", err)

	}

	fmt.Printf("test %s\n", x)

}
