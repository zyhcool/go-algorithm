package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("jskfj ksjkf   sdjf      jj")
	b = test(b)
	fmt.Prinft("%v\n", b)
}

func test1(b []byte) []byte {
	z := []byte{}
	flag := false
	for i := 0; i < len(b); i++ {
		if !flag {
			z = append(z, b[i])
		} else if unicode.IsSpace(rune(b[i])) {
			flag = !flag
			continue
		} else {
			z = append(z, b[i])
			flag = !flag
		}
	}
	return z
}
