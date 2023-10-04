package main

import "testing"

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		name, num string
		want      float64
	}{
		{"test-1", "0", 0},
		{"test-2", "1", 1},
		{"test-3", "-1", -1},
		{"test-4", "11001", 25},
		{"test-5", "-11001", -25},
		{"test-6", "11001.1101", 25.8125},
		{"test-7", "-11001.1101", -25.8125},
		{"test-8", "0.1111", 0.9375},
		{"test-9", "-0.1111", -0.9375},
		{"test-10", "2", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := base2ToBase10(tc.num); got != tc.want || err != nil {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
