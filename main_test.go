package main

import "testing"

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		in  string
		out float64
	}{
		{"11.11", 3.75},
		{"1000.0001", 8.0625},
		{"0.1", 0.5},
		{"1", 1},
		{"100", 4},
		{"101", 5},
		{"110", 6},
		{"111", 7},
		{"1000", 8},
	} {
		if res, err := base2ToBase10(tc.in); err != nil || res != tc.out {
			t.Errorf("base2ToBase10(%v) Expected %v, got %v, %v", tc.in, tc.out, res, err)
		}
	}
}
