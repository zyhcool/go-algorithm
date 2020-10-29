package main

import "fmt"

type point struct {
	x int
	Y int
}

type circle struct {
	point
	Radius int
}

type Wheel struct {
	circle
	Spokes int
}

func main() {
	w := Wheel{circle{point{2, 3}, 10}, 20}
	fmt.Println(w.x)
	fmt.Println(w.circle.point.x)
}
