package main

import (
	"strings"
	"tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	arr := strings.Fields(s)
	for _, val := range arr {
		ret[val]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
