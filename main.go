package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

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
	// Get whole and decimal parts
	// Divide whole part by 2 until it's 0
	// Add the residue to the slice straight away

	// Whole part
	for wp := math.Floor(num); wp != 0; {
		wp = wp / 2
		if wp-math.Floor(wp) > 0 {
			out = "1" + out
		} else {
			out = "0" + out
		}
		wp = math.Floor(wp)
	}
	return out
}

func main() {
	a := base10ToBase2(22.0)
	fmt.Println(a)
}
