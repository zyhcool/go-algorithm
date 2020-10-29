package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	slc := arr[0:2]
	fmt.Printf("array: %v\n", arr)
	fmt.Printf("slice: %v\n", slc)
	fmt.Printf("length of slice: %d\n", len(slc))
	fmt.Printf("capacity of slice: %d\n", cap(slc))
	// fmt.Printf("slice[3]: %d\n", slc[2]) // 编译错误：index out of range [2] with length 2

	reverse(slc)
	fmt.Printf("slice after slice reverse: %v\n", slc)
	fmt.Printf("array after slice reverse: %v\n", arr)

	fmt.Printf("array after array reverse: %v\n", arr)
	reverseArray(arr)
	fmt.Printf("array after array reverse: %v\n", arr)
	// slice是array的引用类型，

	slc2 := []int{1, 2, 3, 4}
	slc2 = append(slc2, 5)
	fmt.Printf("slice2: %v\n", slc2)

	array2 := [4]int{1, 2, 3, 4}
	// array2 = append(array2, 5) // 编译错误：first argument to append must be slice; have [4]int
	fmt.Printf("array2: %v\n", array2)

	slc = append(slc, 5)
	slc = append(slc, 6)
	fmt.Printf("slice after append: %v\n", slc)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseArray(s [4]int) [4]int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
