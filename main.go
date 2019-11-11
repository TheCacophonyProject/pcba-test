package main

import (
	"log"

	"github.com/alexflint/go-arg"
)

var (
	version = "<not set>"
)

type Args struct {
}

func (Args) Version() string {
	return version
}

func procArgs() Args {
	args := Args{}
	arg.MustParse(&args)
	return args
}

func main() {
	err := runMain()
	if err != nil {
		log.Fatal(err)
	}
}

type Tests struct {
	Passed []string
	Failed []string
}

func (t *Tests) addFail(message string) {
	t.Failed = append(t.Failed, message)
}
func (t *Tests) addPass(message string) {
	t.Passed = append(t.Passed, message)
}

func runMain() error {
	_ = procArgs()

	t := Tests{}

	log.Println("Checking RTC:")

	testRTC(&t)

	if len(t.Failed) == 0 {
		log.Println("all tests passed")
	} else {
		log.Println("some test failed")
		log.Printf("failed tests: %v", t.Failed)
	}

	return nil
}
