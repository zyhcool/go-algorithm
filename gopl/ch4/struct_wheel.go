package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	wheel := Wheel{Circle: Circle{Point: Point{X: 1, Y: 2}, Radius: 10}, Spokes: 18}
	fmt.Println(wheel.X)
	fmt.Println(wheel.Y)
	fmt.Println(wheel.Radius)
	fmt.Println(wheel.Spokes)

	fmt.Println(wheel.Circle.Radius)  // 类似 wheel.Radius
	fmt.Println(wheel.Circle.Point.X) // 类似 wheel.X

}
