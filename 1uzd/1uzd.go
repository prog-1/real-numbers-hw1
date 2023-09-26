package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const base2 = "01"

func isValid(input string) bool {
	l := len(input) - 1
	if input[l] == '.' {
		return false
	}
	for _, r := range input {
		if r != '0' && r != '1' && r != '.' && r != '-' {
			return false
		}
	}
	return true
}

func base2ToBase10(num string) (float64, error) {
	if !isValid(num) {
		return 0.0, errors.New("invalid input")
	}

	sign := 1.0
	var intP, fractionalP string
	var convert2 float64
	if strings.Contains(num, "-") {
		_, a, _ := strings.Cut(num, "-")
		sign = -1
		num = a
	}
	if strings.Contains(num, ".") {
		parts := strings.Split(num, ".")
		if len(parts) != 2 {
			return 0.0, errors.New("invalid input")
		}

		intP, fractionalP = parts[0], parts[1]
		length2 := len(fractionalP) - 1
		for i := length2; i >= 0; i-- {
			number := strings.Index(base2, string(fractionalP[i]))
			convert2 += float64(number) * math.Pow(2, float64(-(length2+1)))
		}
	} else {
		intP = num
	}
	var convert int
	length := len(intP) - 1
	for i := 0; i <= length; i++ {
		number := strings.Index(base2, string(intP[i]))
		convert += number * int(math.Pow(2, float64(length-i)))
	}

	return (float64(convert) + convert2) * sign, nil
}

func main() {
	num := "-11.001"
	fmt.Println(base2ToBase10(num))
}
