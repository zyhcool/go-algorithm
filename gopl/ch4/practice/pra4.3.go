package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	reverse(&arr)
	fmt.Printf("%v\n", arr)
}

func reverse(s *[4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
