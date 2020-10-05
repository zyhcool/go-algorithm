package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	// 判断前缀
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, res.Body)
	// 练习 1.9，打印出
	fmt.Printf("\n状态码：%v\n", res.Status)
}
