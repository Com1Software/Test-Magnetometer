package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

func main() {
	fmt.Println("Test Multiport Serial Controller")
	gpsport := ""
	rcport := ""
	senport:=""
	gtg:=3
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
	for x := 0; x < len(ports); x++ {
		port, err := serial.Open(ports[x], mode)
		if err != nil {
			log.Fatal(err)
		}
		line := ""
		buff := make([]byte, 1)
		pass:=0
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
			if strings.Contains(string(buff[:n]), "\n") {
		     if pass>3{
				port.Close()
				break
			 }else{
		    	break 
			}
		    }
		}
		if len(line) > 3 {
			switch {
			case line[0:3] == "CH1":
				rcport = ports[x]
			
			case line[0:3] == "$GP":
				gpsport = ports[x]
			
	    	case line[0:3] == "DIS":
		    	senport = ports[x]
	
			}
		}

	}
	if len(gpsport) > 0 {
		fmt.Printf("GPS Port %s\n", gpsport)
		gtg--
	} else {
		fmt.Printf("GPS Port Not Found\n")
	}
	if len(rcport) > 0 {
		fmt.Printf("RC Port %s\n", rcport)
        gtg--
		} else {
		fmt.Printf("RC Port Not Found\n")
	}
	if len(senport) > 0 {
		fmt.Printf("Sensor Port %s\n", senport)
        gtg--
		} else {
		fmt.Printf("Sensor Port Not Found\n")
	}
    if gtg==0{
		fmt.Printf("Good to go\n")
		fmt.Printf("All ports found\n")
	
		}else{
			fmt.Printf("Init Failure  %d ports not found\n", gtg)
    
		}


}
