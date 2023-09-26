package main

import "testing"

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		name, num string
		want      float64
	}{
		{"test-1", "1", 1},
		{"test-2", "-1", -1},
		{"test-3", "11.001", 3.125},
		{"test-4", "-11.001", -3.125},
		{"test-5", "111101.1111", 61.9375},
		{"test-6", "-111101.1111", -61.9375},
		{"test-7", "0.111100000111", 0.939208984375},
		{"test-8", "-1000101", -69},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := base2ToBase10(tc.num); got != tc.want || err != nil {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
