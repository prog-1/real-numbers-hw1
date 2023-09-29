package main

import (
	"errors"
	"math"
	"strconv"
)

func base10ToBase2(num float64) (string, error) {
	var b string
	b10ToB2Whole(&b, int(num))
	if x := num - math.Floor(num); x != 0 {
		b += "."
		b10ToB2Fraction(&b, x)
	}
	return b, nil
}

func b10ToB2Whole(b *string, num int) string {
	if num > 1 {
		b10ToB2Whole(b, num/2)
	}
	*b += strconv.Itoa(num % 2)
	return *b
}

func b10ToB2Fraction(b *string, num float64) string {
	for i, x := 0, num*2; x != 0 && i <= 20; i, x = i+1, x*2 {
		if x >= 1 {
			*b += strconv.Itoa(1)
			x--
		} else {
			*b += strconv.Itoa(0)
		}
	}
	return *b
}

//###############################################################

func base2ToBase10(num string) (float64, error) {
	if !isValid(num) {
		return 0, errors.New("Error: input is not a decimal number")
	}
	r := []rune(num)

	//whole part
	var i int
	var rw []rune
	x := r[0]
	for ; i < len(r) && r[i] != '.'; i++ {
		x = r[i]
		rw = append(rw, x)
	}
	w := b2ToB10Whole(rw)

	//fraction
	var fr float64
	if i < len(r) {
		var rfr []rune
		i++
		for ; i < len(r); i++ {
			x := r[i]
			rfr = append(rfr, x)
		}
		fr = b2ToB10Fraction(rfr)
	}

	return float64(w) + fr, nil
}

func b2ToB10Whole(rw []rune) (w int) {
	for i, x := range rw {
		w += int(x-'0') * int(math.Pow(2, float64(len(rw)-i-1)))
	}
	return w
}

func b2ToB10Fraction(rfr []rune) (fr float64) {
	for i, x := range rfr {
		fr += float64(x-'0') * float64(math.Pow(2, float64((i+1)*-1)))
	}
	return fr
}

func isValid(num string) bool {
	for _, x := range num {
		if x != '1' && x != '0' && x != '.' {
			return false
		}
	}
	return true
}
