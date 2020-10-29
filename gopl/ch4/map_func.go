package main

import "fmt"

func main() {
	mymap := make(map[string]int)
	mymap["age"] = 25

	fmt.Println(mymap["age"])
	changeMap(mymap)
	fmt.Println(mymap["age"])

	// fmt.Println(mymap["age"])
	// change2(&mymap)
	// fmt.Println(mymap["age"])

	// fmt.Println(&mymap["age"]) // 报错：cannot take the address of mymap["age"]

}

func changeMap(m map[string]int) {
	m["age"] = 100
}

// func change2(m *map[string]int) {
// 	m["age"] = 200
// }
