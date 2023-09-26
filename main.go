package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Accuracy threshold
var e float64 = 1e-3

func base2ToBase10(num string) (float64, error) {
	if !regexp.MustCompile(`^-?[01]*[.]?[01]+$`).MatchString(num) {
		return 0, errors.New("IncorrectInput")
	}

	parts := regexp.MustCompile(`[.]`).Split(num, 2) // [1;2]
	var output float64
	for i, a := range parts[0] {
		b, err := strconv.ParseFloat(string(a), 64)
		if err != nil {
			continue
		}
		output += b * math.Pow(2, float64(len(parts[0])-i-1))
	}
	if len(parts) == 2 {
		for i, a := range parts[1] {
			b, _ := strconv.ParseFloat(string(a), 64)
			output += b / math.Pow(2, float64(i+1))
		}
	}
	if num[0] == '-' {
		output = -output
	}
	return output, nil
}

func base10ToBase2(num float64) (out string) {
	// Whole part
	for n := math.Floor(math.Abs(num)); n != 0; {
		n = n / 2
		if n-math.Floor(n) > 0 {
			out = "1" + out
		} else {
			out = "0" + out
		}
		n = math.Floor(n)
	}

	// Handling negative values
	if num < 0 {
		out = "-" + out
		num = math.Abs(num)
	}

	// returns whether a ~ b within the threshold e
	nearlyEqual := func(a, b float64) bool {
		return math.Abs(a-b) < e
	}
	if nearlyEqual(num-math.Floor(num), 0) { // Fractional part is abscent
		return out
	}

	// Fractional part
	out += "."

	// TODO: rewrite using O(log(n)) search
	var h []float64 // history
	// Returns whether x is present in h
	repeated := func(x float64) bool {
		for i := range h {
			if nearlyEqual(x, h[i]) {
				return true
			}
		}
		return false
	}
	// Returns whether a > b within the constant threshold e
	greaterEqual := func(a, b float64) bool {
		return a-b > e || nearlyEqual(a, b)
	}
	for n := num - math.Floor(num); n != 0 && !repeated(n); {
		h = append(h, n)
		if n *= 2; greaterEqual(n, 1) {
			out += "1"
			n -= 1.0
		} else {
			out += "0"
		}
	}
	return out
}

func main() {
	fmt.Println(base10ToBase2(-3.125))
}
