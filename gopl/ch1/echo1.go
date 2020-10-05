package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s = os.Args[i]
		sep = " "
		s += sep
	}
	fmt.Printf(s)
}
