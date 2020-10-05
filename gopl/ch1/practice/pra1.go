package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 1.1
	fmt.Println(os.Args[0])
	// 1.2
	for i, value := range os.Args {
		fmt.Println(i, value)
	}
	// 1.3
	start1 := time.Now()
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Printf(time.Since(start1).Milliseconds())

	start2 := time.Now()
	strings.Join(os.Args, " ")
	fmt.Println(time.Since(start2).Milliseconds())
}
