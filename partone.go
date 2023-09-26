package main

import (
	"errors"
	"math"
	"strings"
)

func base2ToBase10(num string) (float64, error) {
	sign := 1.0
	if strings.HasPrefix(num, "-") {
		sign = -1.0
		num = num[1:]
	}

	parts := strings.Split(num, ".")

	if len(parts) > 2 {
		return 0, errors.New("invalid input")
	}

	integerPart := parts[0]
	fractionalPart := ""

	if len(parts) == 2 {
		fractionalPart = parts[1]
	}

	integerValue := 0.0
	for i := len(integerPart) - 1; i >= 0; i-- {
		if integerPart[i] == '1' {
			integerValue += math.Pow(2, float64(len(integerPart)-1-i))
		} else if integerPart[i] != '0' {
			return 0, errors.New("invalid input")
		}
	}

	fractionalValue := 0.0
	for i := 0; i < len(fractionalPart); i++ {
		if fractionalPart[i] == '1' {
			fractionalValue += math.Pow(2, -float64(i+1))
		} else if fractionalPart[i] != '0' {
			return 0, errors.New("invalid input")
		}
	}

	return sign * (integerValue + fractionalValue), nil
}

// func main() {
// 	num := "101101.101"
// 	fmt.Println(base2ToBase10(num))
// }
