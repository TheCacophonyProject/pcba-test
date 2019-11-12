package main

import (
	"encoding/binary"
	"fmt"
	"time"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

const (
	attinyAddress       = 0x04
	batteryVoltageLoReg = 0x20
	batteryVoltageHiReg = 0x21

	// 3 was just a randomly chosen as the number for the attiny to return
	// to indicate its presence.
	magicReturn = 0x03

	// Check for the ATtiny for up to a minute.
	maxConnectAttempts     = 20
	connectAttemptInterval = 3 * time.Second
)

func testAttiny(t *Tests) {

	if _, err := host.Init(); err != nil {
		t.addFail(err.Error())
		return
	}
	bus, err := i2creg.Open("")
	if err != nil {
		t.addFail(err.Error())
		return
	}
	dev := &i2c.Dev{Bus: bus, Addr: attinyAddress}
	if detectATtiny(dev) {
		t.addPass("coudl connect to attiny")
	} else {
		t.addPass("failed to connect to attiny")
		return
	}

	// TODO figure out what values this should be between
	battery, err := readBatteryValue(dev)
	if err != nil {
		t.addFail(fmt.Sprintf("failed to read battery voltage %v", err))
	} else {
		t.addPass(fmt.Sprintf("battery reading %d", battery))
	}

	// TODO check reset button
	// TODO check that attiny can reset raspberry pi
}

func detectATtiny(dev *i2c.Dev) bool {
	b := make([]byte, 1)
	err := dev.Tx(nil, b)
	return err == nil && b[0] == magicReturn
}

func readBatteryValue(dev *i2c.Dev) (uint16, error) {
	l := make([]byte, 1)
	h := make([]byte, 1)
	if err := dev.Tx([]byte{batteryVoltageLoReg}, l); err != nil {
		return 0, err
	}
	if err := dev.Tx([]byte{batteryVoltageHiReg}, h); err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16([]byte{h[0], l[0]}), nil
}
