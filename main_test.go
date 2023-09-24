package main

import (
	"regexp"
	"testing"
)

func isInputCorrect(num string) bool {
	return regexp.MustCompile(`^-?[01]*[.]?[01]+$`).MatchString(num)
}
func TestIsInputCorrect(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"11.001", true},
		{"A119kmad", false},
		{"11..001", false},
		{"-11.001", true},
		{"--11.001", false},
		{"11001", true},
		{"11001.", false},
	} {
		if got := isInputCorrect(tc.input); got != tc.want {
			t.Errorf("isInputCorrect() failed with input \"%s\": got = %t, want = %t", tc.input, got, tc.want)
		}
	}
}

func TestBase2ToBase10(t *testing.T) {
	type Want struct {
		num float64
		err error
	}
	for _, tc := range []struct {
		input string
		want  Want
	}{
		{"11.001", Want{3.125, nil}},
		{"-11.001", Want{-3.125, nil}},
		{".001", Want{0.125, nil}},
	} {
		if gotNum, gotErr := base2ToBase10(tc.input); gotNum != tc.want.num || gotErr != tc.want.err {
			t.Errorf("base2ToBase10() Failed with input \"%s\": got = %f, %s; want = %f, %s", tc.input, gotNum, gotErr.Error(), tc.want.num, tc.want.err.Error())
		}
	}
}
