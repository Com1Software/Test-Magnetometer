package main

import (
    "fmt"
    "math"
    "time"

    "github.com/Com1Software/go-i2c"
)

const (
    RegisterA    = 0
    RegisterB    = 0x01
    RegisterMode = 0x02
    XAxisH       = 0x03
    ZAxisH       = 0x05
    YAxisH       = 0x07
    Declination  = -0.00669
    pi           = 3.14159265359
)

var bus *i2c.I2C

func MagnetometerInit() {
    bus.WriteRegU8(RegisterA, 0x70)
    bus.WriteRegU8(RegisterB, 0xa0)
    bus.WriteRegU8(RegisterMode, 0)
}

func readRawData(addr byte) int {
    high, _ := bus.ReadRegU8(addr)
    low, _ := bus.ReadRegU8(addr + 1)
    value := int(int16(high)<<8 | int16(low))
    if value > 32768 {
        value -= 65536
    }
    return value
}

func main() {
    bus, _ = i2c.NewI2C(0x1e, 1)
    defer bus.Close()
    MagnetometerInit()
    fmt.Println("Reading Heading Angle")
    for {
        x := readRawData(XAxisH)
        z := readRawData(ZAxisH)
        y := readRawData(YAxisH)
        heading := math.Atan2(float64(y), float64(x)) + Declination
        if heading > 2*pi {
            heading -= 2 * pi
        }
        if heading < 0 {
            heading += 2 * pi
        }
        headingAngle := int(heading * 180 / pi)
        fmt.Printf("Heading Angle = %d - x=%d y=%d z=%d°\n", headingAngle,x,y,z)
        time.Sleep(1 * time.Second)
    }
}
