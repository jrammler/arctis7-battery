package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/sstallion/go-hid"
)

func getBattery(d *hid.Device) int {
	_, err := d.Write([]byte{0x06, 0x14})
	if err != nil {
		log.Fatal(err)
	}
	report := make([]byte, 31)
	_, err = d.Read(report)
	if err != nil {
		log.Fatal(err)
	}
	if report[2] != 0x03 {
		log.Fatal("Headset not connected")
	}

	_, err = d.Write([]byte{0x06, 0x18})
	if err != nil {
		log.Fatal(err)
	}
	_, err = d.Read(report)
	if err != nil {
		log.Fatal(err)
	}

	return int(report[2])
}

func main() {
	delay := flag.Int("d", 0, "Delay in seconds between between outputs. A value of 0 outputs only once.")
	flag.Parse()

	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	d, err := hid.OpenFirst(0x1038, 0x12ad)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if *delay > 0 {
		for true {
			fmt.Printf("%d\n", getBattery(d))
			time.Sleep(time.Duration(*delay) * time.Second)
		}
	} else {
		fmt.Printf("%d\n", getBattery(d))
	}
}
