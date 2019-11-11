package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

const (
	USBPowerPin = "GPIO22"
)

func testUSB(t *Tests) {
	setUSBPower(false)
	time.Sleep(time.Second)
	initialUSBCount, err := usbBusCount()
	if err != nil {
		t.addFail(err.Error())
		return
	}
	setUSBPower(true)
	time.Sleep(time.Second)
	secondUSBCount, err := usbBusCount()
	if err != nil {
		t.addFail(err.Error())
		return
	}
	if secondUSBCount == initialUSBCount+1 {
		t.addPass("USB power test passed")
	} else {
		t.addFail("USB power test failed")
	}
}

func usbBusCount() (int, error) {
	out, err := exec.Command("lsusb").Output()
	if err != nil {
		return 0, err
	}
	outStr := string(out)
	lines := strings.Split(outStr, "\n")
	busCount := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "Bus") {
			busCount++
		}
	}
	return busCount, nil
}

func setUSBPower(on bool) error {
	pin := gpioreg.ByName(USBPowerPin)
	if on {
		if err := pin.Out(gpio.High); err != nil {
			return fmt.Errorf("failed to set USB power pin high: %v", err)
		}
	} else {
		if err := pin.Out(gpio.Low); err != nil {
			return fmt.Errorf("failed to set USB power pin low: %v", err)
		}
	}
	return nil
}
