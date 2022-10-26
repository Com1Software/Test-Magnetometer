package main

import (
	"fmt"

	"golang.org/x/exp/io/i2c"
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
