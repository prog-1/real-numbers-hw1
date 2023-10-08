package main

import "testing"

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		name string
		num  float64
		want string
	}{
		{"test-1", 1, "1"},
		{"test-2", -1, "-1"},
		{"test-3", 10.75, "1010.11"},
		{"test-4", -10.75, "-1010.11"},
		{"test-5", 0.939208984375, "0.111100000111"},
		{"test-6", -0.939208984375, "-0.111100000111"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := base10ToBase2(tc.num); got != tc.want || err != nil {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
