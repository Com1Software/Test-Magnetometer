package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

func main() {
	fmt.Println("Test Multiport Serial Controller")

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No Serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	po := false
	for x := 0; x < len(ports); x++ {
		port, err := serial.Open(ports[x], mode)
		if err != nil {
			fmt.Println(err)
			po = false
		} else {
			po = true
		}
		if po {
			line := ""
			buff := make([]byte, 1)
			for {
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
				}
				if n == 0 {
					port.Close()
					break
				}
				line = line + string(buff[:n])
				if len(line) > 100 {
					fmt.Println(line)
					line = ""
				}
				if strings.Contains(string(buff[:n]), "\n") {
					fmt.Println("Hit Return")
				}
			}
		}
	}

}
