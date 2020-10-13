package main

import "fmt"

func main() {
	arr := []string{"12", "123", "1234", "1234567", ""}
	for _, ele := range arr {
		fmt.Printf("%s\n", comma(ele))
	}
}

func comma(s string) string {
	if len(s) > 3 {
		left := s[:len(s)-3]
		s = comma(left) + "," + s[len(s)-3:]
	}
	return s
}
