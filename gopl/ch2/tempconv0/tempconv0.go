package tempconv0

import (
	"fmt"
	"log"
	"os"
)

type C float64
type F float64

func CtoF(c C) F {
	return F(c / 10)
}

func FtoC(f F) C {
	return C(f * 10)
}

func (c C) String() string {
	return fmt.Sprintf("%g&&", c)
}

var cwd string

func init() {
	fmt.Printf("haha %s", cwd)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Printf("haha %s", cwd)
}
