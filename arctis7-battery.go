package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/sstallion/go-hid"
)

func getBattery(d *hid.Device) (int, error) {
	_, err := d.Write([]byte{0x06, 0x14})
	if err != nil {
		return 0, err
	}
	report := make([]byte, 31)
	_, err = d.Read(report)
	if err != nil {
		return 0, err
	}
	if report[2] != 0x03 {
		return 0, errors.New("Headset not connected")
	}

	_, err = d.Write([]byte{0x06, 0x18})
	if err != nil {
		return 0, err
	}
	_, err = d.Read(report)
	if err != nil {
		return 0, err
	}

	return int(report[2]), nil
}

func main() {
	delay := flag.Int("d", 0, "Delay in seconds between between outputs. A value of 0 outputs only once.")
	flag.Parse()

	if err := hid.Init(); err != nil {
		log.Fatalf("Error when initializing HID library: %v", err)
	}

	// Open the device using the VID and PID.
	d, err := hid.OpenFirst(0x1038, 0x12ad)
	if err != nil {
		log.Fatal("Unable to connect to headset receiver. Is it connected?")
	}

	defer func() {
		err = d.Close()
		if err != nil {
			log.Fatalf("Error when closing HID library: %v", err)
		}
	}()

	if *delay > 0 {
		for true {
			value, err := getBattery(d)
			if err != nil {
				fmt.Print("--\n");
			} else {
				fmt.Printf("%d\n", value)
			}
			time.Sleep(time.Duration(*delay) * time.Second)
		}
	} else {
		value, err := getBattery(d)
		if err != nil {
			fmt.Print("--\n")
			log.Print(err)
		} else {
			fmt.Printf("%d\n", value)
		}
	}
}
