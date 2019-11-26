package main

import (
	"log"

	pcbatest "github.com/TheCacophonyProject/pcba-test"
	"github.com/alexflint/go-arg"
)

var (
	version = "<not set>"
)

type Args struct {
	USBWaitTime int `arg:"--usb-power-wait" help:"time to wait in seconds after powering on usb"`
	RTCAttempts int `arg:"--attiny-attempts" help:"number of attempts when talking to RTC"`
}

func (Args) Version() string {
	return version
}

func procArgs() Args {
	args := Args{
		USBWaitTime: 1,
		RTCAttempts: 1,
	}
	arg.MustParse(&args)
	return args
}

func main() {
	err := runMain()
	if err != nil {
		log.Fatal(err)
	}
}

func runMain() error {
	log.SetFlags(0)
	args := procArgs()

	t := pcbatest.Tests{}

	log.Println("testing RTC")
	pcbatest.TestRTC(args.RTCAttempts, &t)
	log.Println("testing ATtiny")
	pcbatest.TestAttiny(&t)
	log.Println("testing USB")
	pcbatest.TestUSB(args.USBWaitTime, &t)
	log.Println("testing speakers")
	pcbatest.TestSpeakers(&t)

	//TODO Thermal camera test. Might want just to use managementd and thermal-recorder to display a thermal video for this.
	//TODO rs485 test

	log.Println(t)
	return nil
}
