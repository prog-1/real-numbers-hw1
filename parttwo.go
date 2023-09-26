package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func base10ToBase2(num float64) (string, error) {

	if math.IsNaN(num) || math.IsInf(num, 0) {
		return "", errors.New("invalid input")
	}

	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	integerPart := int64(num)
	integerBinary := strconv.FormatInt(integerPart, 2)

	fractionalPart := num - float64(integerPart)
	fractionalBinary := ""
	precision := 52

	for i := 0; i < precision; i++ {
		fractionalPart *= 2
		bit := '0'
		if fractionalPart >= 1.0 {
			bit = '1'
			fractionalPart -= 1.0
		}
		fractionalBinary += string(bit)
	}

	binaryStr := sign + integerBinary
	if len(fractionalBinary) > 0 {
		binaryStr += "." + strings.TrimRight(fractionalBinary, "0")
	}
	return binaryStr, nil
}

func main() {
	fmt.Println(base10ToBase2(13))
}
