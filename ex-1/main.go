package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

func base2ToBase10(num string) (float64, error) {
	iDot, err := checkBase2(num)
	if err != nil {
		return 0, err
	}
	var pow, res float64
	if iDot == -1 {
		pow = float64(len(num) - 1)
	} else {
		pow = float64(len(num[:iDot]) - 1)
	}
	for _, v := range num {
		if v == '.' {
			continue
		}
		if v != '0' {
			res += math.Pow(2.0, pow)
		}
		pow--
	}
	return res, err
}

func base10ToBase2(num float64) string {
	if num == 0 {
		return "0"
	}
	var res string
	n, frac, err := split(strconv.FormatFloat(num, 'f', 8, 64))
	if err != nil {
		log.Fatal(err)
	}
	for n != 0 {
		res += strconv.Itoa(n % 2)
		n /= 2
	}
	res = reverse(res)
	var p []float64
	if frac != 0 {
		res += "."
	}
	lenI := len(res)
	for frac != 0 {
		frac *= 2
		for i, v := range p {
			if v == frac {
				i += lenI
				res = res[:i] + "(" + res[i:]
				res += ")"
				return res
			}
		}
		p = append(p, frac)
		if frac >= 1 {
			res += strconv.Itoa(1)
			frac, err = strconv.ParseFloat(strconv.FormatFloat(frac, 'f', 8, 64)[1:], 64)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			res += strconv.Itoa(0)
		}
	}
	return res
}

func split(num string) (int, float64, error) {
	iDot := -1
	for i, v := range num {
		if v == '.' {
			iDot = i
			break
		}
	}
	integer, err1 := strconv.Atoi(string([]rune(num)[:iDot]))
	if err1 != nil {
		return 0, 0, err1
	}
	frac, err2 := strconv.ParseFloat(string([]rune(num)[iDot:]), 64)
	if err2 != nil {
		return 0, 0, err2
	}
	return integer, frac, nil
}

func reverse(s string) (res string) {
	for _, v := range s {
		res = string(v) + res
	}
	return
}

func checkBase2(num string) (int, error) {
	iDot := -1
	if num == "" {
		return iDot, errors.New("invalid input")
	}
	for i, v := range num {
		if v == '.' {
			if iDot != -1 {
				return iDot, errors.New("invalid input")
			}
			if v == []rune(num)[len([]rune(num))-1] {
				return iDot, errors.New("invalid input")
			}
			iDot = i
		} else if v != '0' && v != '1' {
			return iDot, errors.New("invalid input")
		}
	}
	return iDot, nil
}

func main() {
	fmt.Println(base10ToBase2(5.7))
}
