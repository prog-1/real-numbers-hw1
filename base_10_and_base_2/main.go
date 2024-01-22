package main

import (
	"errors"
	"fmt"
	"math"
)

func base2ToBase10(num string) (float64, error) {
	var res float64
	pt, sign := -1, 1
	if num[0] == '-' {
		num, sign = num[1:], -1
	}
	for i := range num { // check if the input is correct and find the index of the point
		if (num[i] != '0' && num[i] != '1' && num[i] != '.') || (num[i] == '.' && pt != -1) {
			return 0, errors.New("invalid input")
		}
		if num[i] == '.' {
			pt = i
		}
	}
	intPart := num // convert the integer part
	if pt != -1 {
		intPart = num[:pt]
	}
	for i, l := 0, len(intPart)-1; i < len(intPart); i, l = i+1, l-1 {
		if intPart[i] == '1' {
			res += math.Pow(2, float64(l))
		}
	}
	if pt != -1 { // convert the fractional part
		fracPart := num[pt+1:]
		for i := range fracPart {
			if fracPart[i] == '1' {
				res += math.Pow(2, float64(-i-1))
			}
		}
	}
	return res * float64(sign), nil
}

func roundFloat(num float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(num*ratio) / ratio
}

func base10ToBase2(num float64) (string, error) {
	var res string
	if num < 0 {
		res += "-"
		num *= -1
	}
	if num >= 1 { // convert the integer part
		i, currNum := 1, int(math.Floor(num))
		for ; i <= currNum; i *= 2 {
		}
		for i /= 2; i >= 1; i /= 2 {
			res += fmt.Sprint(currNum / i)
			if currNum >= i {
				currNum -= i
			}
		}
		num -= math.Floor(num)
	} else {
		res += "0"
	}
	if num > 0 { // convert the fractional part
		res += "."
		pt := len(res)
		precision := len(fmt.Sprint(num)) - 2
		for x := num * 2; ; x *= 2 {
			if x >= 1 {
				res += "1"
				x -= 1
			} else {
				res += "0"
			}
			x = roundFloat(x, precision) // otherwise there are problems
			if x == 0 {
				break
			}
			if x == num {
				res = res[:pt] + "(" + res[pt:] + ")"
				break
			}
		}
	}
	return res, nil
}

func main() {
	fmt.Println(base2ToBase10("101101.101"))
	fmt.Println(base10ToBase2(0.2))
	fmt.Println(base10ToBase2(45.625))
}
