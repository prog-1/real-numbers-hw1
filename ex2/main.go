package main

import (
	"fmt"
	"math"
)

func base10ToBase2(num float64) (string, error) {
	if num == 0 {
		return "0", nil
	}

	intPart := int(math.Abs(num))
	fracPart := math.Abs(num) - float64(intPart)
	intBinary := decimalToBinary(intPart)
	fracBinary := decimalFractionToBinary(fracPart)

	var res string
	if num < 0 {
		res += "-"
	}
	if math.Abs(num) < 1 {
		res += "0"
	}
	if fracBinary != "" {
		res = res + intBinary + "." + fracBinary
	} else {
		fmt.Println(res)

		res = res + intBinary
	}

	return res, nil
}

func decimalToBinary(num int) string {
	var res string
	for num > 0 {
		res = fmt.Sprintf("%d", num%2) + res
		num /= 2
	}
	return res
}

func decimalFractionToBinary(frac float64) string {
	var res string
	for frac > 0 {
		frac *= 2
		whole := int(frac)
		res += fmt.Sprintf("%d", int(frac))
		frac -= float64(whole)
	}

	return res
}
