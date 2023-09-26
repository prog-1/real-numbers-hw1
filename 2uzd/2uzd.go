package main

import (
	"errors"
	"fmt"
	"math"
)

const base2 = "01"

func base10ToBase2(num float64) (string, error) {
	if math.IsNaN(num) || math.IsInf(num, 0) {
		return "", errors.New("invalid input")
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	intP := int64(num)
	fractionalP := num - float64(intP)

	var z int64
	var number float64
	var convert string
	for number <= float64(intP) {
		z++
		number = math.Pow(2, float64(z))
	}
	number /= float64(2)
	for z > 0 {
		z--
		convert += string(base2[intP/int64(number)])
		intP %= int64(number)
		number /= 2
	}

	var convert2 string
	for i := 0; i < 64; i++ {
		fractionalP *= 2
		bit := int(fractionalP)
		convert2 += fmt.Sprintf("%d", bit)
		fractionalP -= float64(bit)
		if fractionalP == 0 {
			break
		}
	}

	res := convert
	if convert2 != "" {
		res += "." + convert2
	}

	return sign + res, nil
}

func main() {
	num := 3.125
	fmt.Println(base10ToBase2(num))
}
