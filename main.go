package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func base2ToBase10(num string) (float64, error) {
	// Split num in two parts
	a := strings.Split(num, ".")
	if len(a) != 2 && len(a) != 1 {
		return 0, errors.New("Invalid number")
	}
	a0, err := strconv.Atoi(a[0])
	if err != nil {
		return 0, err
	}
	var res float64
	for i := 0; a0 > 0; a0, i = a0/10, i+1 {
		res += float64(a0%10) * math.Pow(2, float64(i))
	}
	if len(a) == 1 {
		return res, nil
	}
	a1, err := strconv.Atoi(a[1])
	if err != nil {
		return res, err
	}
	for i := -len(a[1]); a1 > 0; a1, i = a1/10, i+1 {
		res += float64(a1%10) * math.Pow(2, float64(i))
	}
	return res, nil

}

func main() {
	fmt.Println(base2ToBase10("1000"))
}
