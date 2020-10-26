package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["name"] = 12
	m["age"] = 99
	for k, v := range m {
		fmt.Printf("key:%s, value:%d\n", k, v)
	}
	value, ok := m["zyh"]
	if !ok {
		fmt.Printf("no key:zyh\t%v", value)
	}
}
