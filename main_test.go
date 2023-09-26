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
	for _, tc := range []struct {
		input string
		want  float64
	}{
		{"11", 3},
		{"11.001", 3.125},
		{"-11.001", -3.125},
		{".001", 0.125},
	} {
		if got, _ := base2ToBase10(tc.input); got != tc.want {
			t.Errorf("base2ToBase10() failed with input \"%s\": got = %f, want = %f", tc.input, got, tc.want)
		}
	}
}

func TestBase10ToBase2(t *testing.T) {
	for _, tc := range []struct {
		input float64
		want  string
	}{
		{22.0, "10110"},
		{3.125, "11.001"},
		{-3.125, "-11.001"},
		{0.125, ".001"},
	} {
		if got := base10ToBase2(tc.input); got != tc.want {
			t.Errorf("base10ToBase10() failed with input \"%f\": got = %s, want = %s", tc.input, got, tc.want)
		}
	}
}
