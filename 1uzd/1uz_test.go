package main

import (
	"errors"
	"testing"
)

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want float64
		err  error
	}{
		{num: "11.001", want: 3.125, err: nil},
		{num: "-11.001", want: -3.125, err: nil},
		{num: "1011.1001", want: 11.5625, err: nil},
		{num: "111101", want: 61, err: nil},
		{num: "1", want: 1, err: nil},
		{num: "0", want: 0, err: nil},
		{num: "1.", want: 0, err: errors.New("invalid input")},
		{num: "-1.", want: 0, err: errors.New("invalid input")},
		{num: "111111000002111.1010", want: 0, err: errors.New("invalid input")},
		{num: "e11011000111.1010", want: 0, err: errors.New("invalid input")},
		{num: "11011000111.1A010", want: 0, err: errors.New("invalid input")},
		{num: "11011000111.10.10", want: 0, err: errors.New("invalid input")},
		{num: "-110110.00111.10", want: 0, err: errors.New("invalid input")},
		{num: "-z10110.0011110", want: 0, err: errors.New("invalid input")},
		{num: "110O", want: 0, err: errors.New("invalid input")},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got, got2 := base2ToBase10(tc.num); got != tc.want && got2 != tc.err {
				t.Errorf("got = %v, %v, want = %v", got, got2, tc.want)
			}
		})
	}
}
