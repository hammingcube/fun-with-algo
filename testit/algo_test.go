package algo

import (
	"fmt"
	"gopkg.in/pipe.v2"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	p1 := pipe.Line(
		pipe.ChDir("../solutions/prob1/v1/"),
		pipe.ReadFile("input.txt"),
		pipe.Exec("./runit"),
	)
	p2 := pipe.Line(
		pipe.ReadFile("../problems/prob1/v1/output.txt"),
	)
	got, err := pipe.CombinedOutput(p1)
	expected, err := pipe.CombinedOutput(p2)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("got=%s expected=%s\n", strings.TrimSpace(string(got)), strings.TrimSpace(string(expected)))
}
