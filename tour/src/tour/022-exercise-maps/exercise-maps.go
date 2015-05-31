package main 

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	
	for _, v := range strings.Fields(s) {
		_, has := m[v]
		if has {
			m[v]++;
		} else {
			m[v] = 1
		}
	}
	
	return
}

func main() {
	wc.Test(WordCount)
}

