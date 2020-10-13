package main

import (
	"bytes"
	"fmt"
)

func main() {
	arr := []string{"12", "123", "1234", "1234567", ""}
	for _, ele := range arr {
		fmt.Printf("%s\n", comma(ele))
	}
}

func comma(s string) string {
	start := len(s) % 3
	n := int(len(s) / 3)
	var res bytes.Buffer
	if start != 0 {
		res.WriteString(s[0:start])
	}
	for ; n > 0; n-- {
		if start <
		res.WriteString(",")
		res.WriteString(s[start-3 : start])
		start += 3
	}
	return res.String()
}
