package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(string(0x4e16))
	for a, b := range s {
		fmt.Println(a, " ", b)
	}

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i, r, size)
		i += size
	}
	fmt.Printf("%q", s[4])
}
