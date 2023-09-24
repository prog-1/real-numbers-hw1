package main

import (
	"errors"
	"testing"
)

func TestBase2ToBase10(t *testing.T) {
	for _, tc := range []struct {
		name string
		num  string
		want float64
		err  error
	}{
		{"case-1", "11", 3, nil},
		{"case-2", "101", 5, nil},
		{"case-3", "11.001", 3.125, nil},
		{"case-4", "11.11", 3.75, nil},
		{"case-5", "101.101", 5.625, nil},
		{"case-6", ".01", 0.25, nil},
		{"case-7", "1.", 0, errors.New("invalid input")},
		{"case-8", "1..1", 0, errors.New("invalid input")},
		{"case-9", "2", 0, errors.New("invalid input")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := base2ToBase10(tc.num)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
			if err != nil || tc.err != nil {
				if err.Error() != tc.err.Error() {
					t.Errorf("err = %v, want_err = %v", err, tc.err)
				}
			}
		})
	}
}

func TestBase10ToBase2(t *testing.T) {
	for _, tc := range []struct {
		name string
		num  float64
		want string
	}{
		{"case-1", 3, "11"},
		{"case-2", 5, "101"},
		{"case-3", 3.125, "11.001"},
		{"case-4", 3.75, "11.11"},
		{"case-5", 5.625, "101.101"},
		{"case-6", 0.25, ".01"},
		{"case-7", 0.2, ".(0011)"},
		{"case-8", 5.7, "101.1(0110)"},
		{"case-9", 0, "0"},
		{"case-10", 0.0, "0"},
		{"case-11", .0, "0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := base10ToBase2(tc.num)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
