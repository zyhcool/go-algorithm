package main

import "fmt"

func main() {
	// s := []string{"1", "1", "3", "4", "4", "4", "5"}
	s := []string{"1"}
	s = test(s)
	fmt.Printf("%v\n", s)
}

func test(s []string) []string {
	if len(s) == 0 {
		return s
	}
	z := s[:1]
	for i, j := 0, 1; j < len(s); i, j = i+1, j+1 {
		if s[i] == s[j] {
			continue
		} else {
			z = append(z, s[j])
		}
	}
	return z
}
