package main

import (
	"errors"
	"math"
	"strings"
)

func base2ToBase10(num string) (float64, error) {
	parts := strings.Split(num, ".")
	if len(parts) > 2 || parts[0] == "" {
		return 0, errors.New("invalid character")
	}
	if len(parts) == 1 {
		ItnBase, err := IntegerBase(parts[0])
		if err != nil {
			return 0, err
		}
		return ItnBase, nil
	}
	if len(parts) == 2 {
		FLoatBase, err := FractBase(parts)
		if err != nil {
			return 0, err
		}
		return FLoatBase, nil
	}
	return 0, nil
}

func IntegerBase(part string) (float64, error) {
	var res, starpres, i float64
	if part[0] == 45 {
		part = part[1:]
		res = -1
	} else {
		res = 1.0
	}
	i = float64(len(part) - 1)
	for _, v := range part {
		if v == '1' {
			starpres += math.Pow(2, i)
		} else if v != '0' {
			return 0, errors.New("invalid character")
		}
		i--
	}
	starpres *= res
	return starpres, nil
}

func FractBase(parts []string) (float64, error) {
	var starpres, i float64
	res, err := IntegerBase(parts[0])
	if err != nil {
		return 0, err
	}
	i = -1
	for _, v := range parts[1] {
		if v == '1' {
			starpres += math.Pow(2, i)
		} else if v != '0' {
			return 0, errors.New("invalid character")
		}
		i--
	}
	res += starpres
	return res, nil
}
