package main

import "fmt"

func main() {
	type Person struct {
		name    string
		age     int
		country string
		hobbies []string
	}

	zyh := Person{
		name:    "zyh",
		age:     26,
		country: "China",
		hobbies: []string{"basketball"},
	}
	fmt.Println(zyh)

	name := &zyh.name
	fmt.Println(*name)

	pzyh := &zyh
	age := pzyh.age
	fmt.Println(age)

	page := &pzyh.age
	*page = 99
	fmt.Println(zyh.age)
}
