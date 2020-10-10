package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "zyh", "the name of the true king")

func main() {
	flag.Parse()
	fmt.Printf("hello ,the true king is %s", *name)
}
