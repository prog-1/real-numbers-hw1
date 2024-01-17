package main

import (
	"fmt"
	"math"
)

func PatternMatching(s [][]string) {
	var res []string
	for _, x := range s {
		currRes := x[0]
		for i := 1; i < len(x); i++ {
			current := x[i]
			if currRes[0] == '*' {
				diff := int(math.Abs(float64(len(currRes) - len(current))))
				if current[0] == '*' {
					if len(currRes) < len(current) {
						if currRes[1:] == current[diff+1:] {
							currRes = current
						} else {
							currRes = "*"
							break
						}
					} else if len(currRes) >= len(current) && currRes[diff+1:] != current[1:] {
						currRes = "*"
						break
					}
				}
			}
		}
		if currRes[0] == '*' && len(currRes) > 1 {
			currRes = currRes[1:]
		}
		res = append(res, currRes)
	}
	for i := range res {
		fmt.Printf("Case #%v: %v\n", i+1, res[i])
	}
}

func main() {
	PatternMatching([][]string{{"*CONUTS", "*COCONUTS", "*OCONUTS", "*CONUTS", "*S"}, {"*XZ", "*ZYX"}, {"H*O", "HELLO*", "*HELLO", "HE*"}})
}
