package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	// usage: go run 02_echo.go hello world
	echo1()
	echo2()
}
