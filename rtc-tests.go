package main

import (
	"fmt"
	"time"

	"github.com/TheCacophonyProject/rtc-utils/rtc"
)

func testRTC(t *Tests) {

	// RTC Battery. If the battery is low the test will also fail
	// Battery test is repeated as if the battery plug is not connected it can be
	// a floating voltage on the RTC battery pin
	batteryPassed := true
	for i := 0; i < 20; i++ {
		if rtc.CheckBattery(1) != nil {
			batteryPassed = false
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if batteryPassed {
		t.addPass("RTC battery test passed")
	} else {
		t.addFail("RTC battery test failed. Check that battery is plugged in HAT and isn't flat")
	}

	now := time.Now().UTC() // Used for checking read test
	if rtc.Write(1) == nil {
		t.addPass("RTC time write passed")
	} else {
		t.addFail("RTC time write failed")
		return // Don't test read if write failed
	}

	state, err := rtc.State(1)
	if err != nil {
		t.addFail("RTC state read failed")
		return
	}
	if state.Time.Sub(now) > time.Second {
		t.addFail(fmt.Sprintf("RTC write or read time failed. Time written %s, time read %s", now.UTC(), state.Time))
	} else {
		t.addPass("RTC read and write failed")
	}
}
