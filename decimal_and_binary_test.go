package main

import (
	"testing"
)

func TestBase10ToBase2(t *testing.T) {
	for _, tc := range []struct {
		input float64
		want  string
	}{
		{3.125, "11.001"},
		{10.1, "1010.000110011001100110011"},
		{7, "111"},
		{0.2, "0.001100110011001100110"},
		{100.0, "1100100"},
		{0, "0"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			got, _ := base10ToBase2(tc.input)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  float64
	}{
		{"11.001", 3.125},
		{"10.01", 2.25},
		{"111", 7},
		{"1.01", 1.25},
		{"as.bb", 0},
		{"435.234", 0},
		{"110.10f0", 0},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, _ := base2ToBase10(tc.input)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Base10ToBase2", F: TestBase10ToBase2},
			{Name: "Test Base2ToBase10", F: TestBase2ToBase10},
		},
		/* benchmarks */ nil /* examples */, nil)
}
