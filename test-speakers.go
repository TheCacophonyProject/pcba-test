package pcbatest

import (
	"fmt"
	"os/exec"
)

func TestSpeakers(t *Tests) {
	//speaker-test --test=wav -w /usr/share/sounds/alsa/Front_Center.wav -l 1

	_, err := exec.Command(
		"speaker-test",
		"--nloops=1",
		"--test=wav",
		"--wavfile=/usr/share/sounds/alsa/Front_Center.wav",
	).Output()
	if err != nil {
		t.addFail(fmt.Sprintf("error with playing audio: %v", err))
	}
}
