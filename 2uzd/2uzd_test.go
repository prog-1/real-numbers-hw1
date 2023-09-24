package main

import (
	"testing"
)

func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		num  float64
		want string
		err  error
	}{
		{num: 10, want: "1010", err: nil},
		{num: 0, want: "0", err: nil},
		{num: 0.01, want: "0.00000010100", err: nil},
		{num: -0.01, want: "-0.00000010100", err: nil},
		{num: 13614146.12634346, want: "110011111011110001000010.00100000010", err: nil},
		{num: 1, want: "1", err: nil},
		{num: -1, want: "-1", err: nil},
		{num: 3.125, want: "11.001", err: nil},
		{num: -3.125, want: "-11.001", err: nil},
	} {
		got, got2 := base10ToBase2(tc.num)
		if got != tc.want && got2 != tc.err {
			t.Errorf("got: %v,%v, want: %v", got, got2, tc.want)
		}
	}
}
