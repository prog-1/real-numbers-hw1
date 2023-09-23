package main

import (
	"errors"
	"math"
	"strings"
)

func base2ToBase10(num string) (float64, error) {
	var res float64
	if num[0] == 45 {
		num = num[1:]
		res = -1
	} else {
		res = 1.0
	}
	parts := strings.Split(num, ".")
	if len(parts) < 2 {
		intPart, err := binaryIntToDecimal(parts[0])
		if err != nil {
			return 0, err
		}
		return res * intPart, nil
	}
	integer := parts[0]
	fraction := parts[1]

	intPart, err := binaryIntToDecimal(integer)
	if err != nil {
		return 0, err
	}

	fracPart, err := binaryFractionToDecimal(fraction)
	if err != nil {
		return 0, err
	}

	return res * (intPart + fracPart), nil
}

func binaryIntToDecimal(binaryStr string) (float64, error) {
	var res, pow float64
	pow = float64(len(binaryStr) - 1)
	for _, number := range binaryStr {
		if number == '1' {
			res += math.Pow(2, pow)

		} else if number != '0' {
			return 0, errors.New("Invalid character in binary string")
		}
		pow--
	}

	return res, nil
}

func binaryFractionToDecimal(binaryStr string) (float64, error) {
	var res, pow float64
	pow = -1.0
	for _, number := range binaryStr {
		if number == '1' {
			res += math.Pow(2, pow)
		} else if number != '0' {
			return 0, errors.New("Invalid character in binary string")
		}
		pow--
	}

	return res, nil
}
