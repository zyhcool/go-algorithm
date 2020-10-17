package main

import "fmt"

func main() {
	var s = []int{1, 2}
	s = appendInt(s, 4)
	fmt.Printf("%v\n", s)
}

func appendInt(s []int, e int) []int {
	z := []int{}
	zlen := len(s) + 1
	if zlen <= cap(s) {
		z = s
	} else {
		zcap := zlen
		if zcap < len(s)*2 {
			zcap = len(s) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, s)
	}
	z[zlen-1] = e
	return z

}
