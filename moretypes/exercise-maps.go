package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, value := range strings.Fields(s) {
		result[value] += 1
	}

	return result
}
func main() {
	wc.Test(WordCount)
}
