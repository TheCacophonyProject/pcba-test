package pcbatest

import "fmt"

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
