package main

import "fmt"

type Person struct {
	name string
}

func main() {
	zyh := Person{
		name: "zyh",
	}
	fmt.Println(zyh.name)
	change(zyh)
	fmt.Println(zyh.name)
	fmt.Println(zyh.name)
	change2(&zyh)
	fmt.Println(zyh.name)
}

func change(p Person) {
	p.name = "haha"
}

func change2(p *Person) {
	p.name = "xixi"
}
