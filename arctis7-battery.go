package main

import (
	"os"
	"fmt"
	"log"

	"github.com/sstallion/go-hid"
)

func main() {
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	d, err := hid.OpenFirst(0x1038, 0x12ad)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}

	_, err = d.Write([]byte{0x06, 0x14})
	if err != nil {
		log.Fatal(err)
	}
	report := make([]byte, 31)
	_, err = d.Read(report)
	if err != nil {
		log.Fatal(err)
	}
	if report[2] != 0x03 {
		log.Printf("Headset not connected\n")
		os.Exit(3)
	}

	_, err = d.Write([]byte{0x06, 0x18})
	if err != nil {
		log.Fatal(err)
	}
	_, err = d.Read(report)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", report[2])
}
