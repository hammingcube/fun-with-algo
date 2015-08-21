package algo

import (
	"fmt"
	"gopkg.in/pipe.v2"
	"testing"
)

func TestAverage(t *testing.T) {
	p1 := pipe.Line(
		pipe.ReadFile("../solutions/input.txt"),
	)
	got, err := pipe.CombinedOutput(p1)
	if err != nil {
		t.Error("Something went wrong!")
	} else {
		fmt.Printf("%s", got)
	}
}
