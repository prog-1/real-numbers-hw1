package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	// num := "11.11"
	// out, err := base2ToBase10(num)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(out)
	num10 := 5.7
	fmt.Println(base10ToBase2(num10))
}

func base2ToBase10(num string) (float64, error) {
	var res float64
	isDec, dotI, err := check2ndBase(num)
	if err != nil {
		return 0, err
	}

	if !isDec {
		pow := float64(len(num) - 1)
		for _, j := range num {
			if j == '1' {
				res += math.Pow(2, pow)
			}
			pow--
		}
	} else {
		pow := float64(len(num[:dotI]) - 1)
		for i, j := range num {
			if i == dotI {
				continue
			}
			if j == '1' {
				fmt.Println(res)
				res += math.Pow(2, pow)
			}
			pow--
		}
	}
	return res, nil
}

func check2ndBase(num string) (bool, int, error) {
	isDec := false
	var dotI int
	if num == "" {
		return false, dotI, errors.New("invalid input")
	}

	for i, j := range num {
		if j == '.' {
			dotI = i
			if isDec || i+1 == len(num) {
				return false, dotI, errors.New("invalid input")
			}
			isDec = true
		} else if j != '0' && j != '1' {
			return false, dotI, errors.New("invalid input")
		}
	}
	return isDec, dotI, nil
}

func base10ToBase2(num float64) string {
	if num == 0 {
		return "0"
	}

	var resfull, resfrac string
	full, frac := split(num)

	for full != 0 {
		resfull += strconv.Itoa(full % 2)
		full = full / 2
	}
	resfull = reverse(resfull)

	periodI := -1
	var v []float64
	var f []string

	if frac != 0 {
		for frac != 0 {
			frac *= 2
			fmt.Println(f)
			fmt.Println(v)
			fmt.Println(frac, "----------")
			for i, j := range v {
				if frac == j {
					periodI = i
					break
				}
			}
			if periodI != -1 {
				break
			}
			v = append(v, frac)
			if frac >= 1 {
				f = append(f, "1")
				frac, _ = strconv.ParseFloat(strconv.FormatFloat(frac, 'f', 8, 64)[1:], 64)
			} else {
				f = append(f, "0")
			}
		}

		if periodI != -1 {
			fmt.Println(f)
			resfrac = "."
			f = append(f, ")")
			for i, j := range f {
				if i == periodI {
					resfrac = resfrac + "(" + j
				} else {
					resfrac += j
				}
			}
		} else {
			resfrac = "."
			for _, j := range f {
				resfrac += j
			}

		}
	}

	return resfull + resfrac
}

func split(num float64) (int, float64) {
	n := strconv.FormatFloat(num, 'f', 12, 64)
	dotI := -1
	for i, j := range n {
		if j == '.' {
			dotI = i
			break
		}
	}
	if dotI != -1 {
		full, _ := strconv.Atoi(n[:dotI])
		frac, _ := strconv.ParseFloat(n[dotI:], 64)
		return full, float64(frac)
	}
	return int(num), 0
}

func reverse(num string) (res string) {
	for _, j := range num {
		res = string(j) + res
	}
	return res
}
