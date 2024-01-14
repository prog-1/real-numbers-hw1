package main

import (
	"testing"
)

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  float64
	}{
		{"1", "0", 0},
		{"2", "1", 1},
		{"3", "-1", -1},
		{"4", "0.1", 0.5},
		{"5", "-0.1", -0.5},
		{"6", "10", 2},
		{"7", "100", 4},
		{"8", "101", 5},
		{"9", "110", 6},
		{"10", "111", 7},
		{"11", "1000", 8},
		{"12", "-1000", -8},
		{"13", "11.001", 3.125},
		{"14", "-11.001", -3.125},
		{"15", "101101.101", 45.625},
		{"16", "-101101.101", -45.625},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, _ := base2ToBase10(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestBase10ToBase2(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input float64
		want  string
	}{
		{"1", 0, "0"},
		{"2", 1, "1"},
		{"3", -1, "-1"},
		{"4", 0.2, "0.(0011)"},
		{"5", -0.2, "-0.(0011)"},
		{"6", 0.5, "0.1"},
		{"7", -0.5, "-0.1"},
		{"8", 2, "10"},
		{"9", 4, "100"},
		{"10", 5, "101"},
		{"11", 6, "110"},
		{"12", 7, "111"},
		{"13", 8, "1000"},
		{"14", -8, "-1000"},
		{"15", 3.125, "11.001"},
		{"16", -3.125, "-11.001"},
		{"17", 45.625, "101101.101"},
		{"18", -45.625, "-101101.101"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, _ := base10ToBase2(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
