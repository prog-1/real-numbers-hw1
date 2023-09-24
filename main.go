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
	for n := math.Floor(num); n != 0; {
		n = n / 2
		if n-math.Floor(n) > 0 {
			out = "1" + out
		} else {
			out = "0" + out
		}
		n = math.Floor(n)
	}
	out += "."

	// Fractional part
	// * 2 and cut the whole part until the fractional part is either 0 or repeats

	// TODO: rewrite using O(log(n)) search
	var h []float64 // history
	// Returns whether a is present in h
	repeated := func(a float64) bool {
		for _, x := range h {
			if x == a {
				return true
			}
		}
		return false
	}

	for n := num - math.Floor(num); n != 0 && !repeated(n); {
		if n = 2 * n; n > 1 {
			out += "1"
			n -= 1.0
		} else {
			out += "0"
		}
	}
	return out
}

func main() {
	a := base10ToBase2(0.2)
	fmt.Println(a)
}

type Heap []int

// Source := https://github.com/34thSchool/heap-1
// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	// Add new element to the end of the Heap
	// compare new element's value with its parent
	// if new element < parent -> swap them
	*h = Heap(append(*h, x))
	for i := len(*h) - 1; i > 0 && (*h)[i] < (*h)[(i-1)/2]; i = (i - 1) / 2 {
		(*h)[i], (*h)[(i-1)/2] = (*h)[(i-1)/2], (*h)[i]
	}
}
