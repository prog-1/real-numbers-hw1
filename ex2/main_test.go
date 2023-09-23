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
		{"test-3", 3.125, "11.001"},
		{"test-4", -3.125, "-11.001"},
		{"test-5", 61.9375, "111101.1111"},
		{"test-6", -61.9375, "-111101.1111"},
		{"test-7", 0.939208984375, "0.111100000111"},
		{"test-8", -69, "-1000101"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := base10ToBase2(tc.num); got != tc.want || err != nil {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
