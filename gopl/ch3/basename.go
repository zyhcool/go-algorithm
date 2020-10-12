package main

import "fmt"

func main() {
	s := "a/b/c.go"
	res := basename(s)
	fmt.Printf("%s", res)
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			fmt.Print(s[i])
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
