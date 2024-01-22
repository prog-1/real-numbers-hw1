package main

import (
	"fmt"
)

func PatternMatching(s [][]string) {
	for num, x := range s {
		currRes := x[0]
		for i := 1; i < len(x); i++ {
			current := x[i]
			if current[0] != '*' {
				if currRes[0] != '*' {
					isWrong := false
					for i := 0; current[i] != '*' && currRes[i] != '*'; i++ {
						if current[i] != currRes[i] {
							isWrong = true
							break
						}
					}
					if isWrong {
						currRes = "*"
						break
					}
					if len(current) > len(currRes) {
						currRes = current
					}
				}
			}
			if current[len(current)-1] != '*' {
				if currRes[len(currRes)-1] != '*' {
					isWrong := false
					for i, j := len(current)-1, len(currRes)-1; current[i] != '*' && currRes[j] != '*'; i, j = i-1, j-1 {
						if current[i] != currRes[j] {
							isWrong = true
							break
						}
					}
					if isWrong {
						currRes = "*"
						break
					}
					if len(current) > len(currRes) {
						currRes = current
					}
				}
			}
		}
		for i := 0; i < len(currRes) && len(currRes) > 1; i++ { // remove all '*'
			if currRes[i] == '*' {
				currRes = currRes[:i] + currRes[i+1:]
			}
		}
		fmt.Printf("Case #%v: %v\n", num+1, currRes)
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	var s [][]string
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		var pat []string
		for ; n > 0; n-- {
			var p string
			fmt.Scan(&p)
			pat = append(pat, p)
		}
		s = append(s, pat)
	}
	PatternMatching(s)
}
