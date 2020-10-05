package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	count := make(map[string]int)
	for input.Scan() {
		count[input.Text()]++
	}

	for key, value := range count {
		if value > 1 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}
