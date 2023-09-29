package main

import (
	"testing"
)

func TestPatternMatching(t *testing.T) {
	for _, tc := range []struct {
		input []string
		want  string
	}{
		{[]string{"*CONUTS", "*COCONUTS", "*OCONUTS"}, "COCONUTS"},
		{[]string{"HE*", "*HOGS", "*DGE*S"}, "HEDGEHOGS"},
		{[]string{"H*O", "HELLO*", "*HELLO", "HE*"}, "HELLOHELLO"},
		{[]string{"CODE*", "*JAM"}, "CODEJAM"},
		{[]string{"CO*DE", "J*AM"}, "*"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			got := patternMatching(tc.input)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func main1() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test TestPatternMatching", F: TestPatternMatching},
		},
		/* benchmarks */ nil /* examples */, nil)
}
