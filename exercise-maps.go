package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)

	counts := make(map[string]int)
	for _, s := range ss {
		_, ok := counts[s]
		if !ok {
			counts[s] = 1
		} else {
			counts[s] = counts[s] + 1
		}

	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
