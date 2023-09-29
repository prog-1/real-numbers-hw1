package main

import (
	"strings"
)

func patternMatching(patterns []string) string {
	var lpre, lpost string

	//saving longest prefix and longest suffix
	for _, p := range patterns {
		var pre, post string
		pre, post = splitPattern(p)
		if len(pre) > len(lpre) { // checking longest prefix saved
			lpre = pre
		}
		if len(post) > len(lpost) { // checking longest postfix saved
			lpost = post
		}
	}

	//checking whether all patterns have longest prefix and suffix. If not return '*'
	for _, p := range patterns {
		if !checkPrefix(lpre, p) || !checkPostfix(lpost, p) {
			return "*"
		}
	}

	//making result string
	r := lpre
	for _, p := range patterns {
		//mid := p[len(lpre):len(lpost)] //taking part without prefix and postfix
		mid := takeMiddle(p)
		if mid != "" {
			r += strings.ReplaceAll(mid, "*", "") //adding to result and deleting all '*'
		}
	}
	r += lpost
	return r
}

func splitPattern(p string) (pre, post string) {
	parts := strings.Split(p, "*")
	pre = parts[0]
	if len(parts) > 1 {
		post = parts[len(parts)-1]
	}
	return pre, post
}

func checkPrefix(pre, p string) bool {
	//p = strings.ReplaceAll(p, "*", "")
	var i, j int
	for ; i < len(pre) && j < len(p)-1 && p[j] != '*'; i++ {
		if pre[i] == p[j] {
			j++
		}
	}
	return j == len(p)-1 || p[j] == '*'
}

func checkPostfix(post, p string) bool {
	//p = strings.ReplaceAll(p, "*", "")
	var i, j int
	for ; i < len(post) && j < len(p)-1 && p[len(p)-1-j] != '*'; i++ {
		if post[len(post)-1-i] == p[len(p)-1-j] {
			j++
		}
	}
	return j == len(p)-1 || p[len(p)-1-j] == '*'
}

func takeMiddle(p string) (mid string) {
	parts := strings.Split(p, "*")
	if len(parts) > 2 {
		parts = parts[1 : len(parts)-1]
		for i := range parts {
			mid += parts[i]
		}
	}
	return mid
}
