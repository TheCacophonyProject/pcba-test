package main

import (
	"fmt"
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

func (t Tests) String() string {
	res := fmt.Sprintf("%d passed, %d failed", len(t.Passed), len(t.Failed))
	res = res + "\npassed tests:"
	for _, s := range t.Passed {
		res = res + "\n\t" + s
	}
	res = res + "\nfailed tests:"
	for _, s := range t.Failed {
		res = res + "\n\t" + s
	}
	return res
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

	log.Println("testing RTC")
	testRTC(&t)
	log.Println("testing ATtiny")
	testAttiny(&t)

	//TODO speaker test
	//TODO USB test
	//TODO Thermal camera test
	//TODO rs485 test

	log.Println(t)
	return nil
}
