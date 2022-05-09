package main

import "fmt"

func main() {
	x := "hello"
	for i := 0; i < len(x); i++ {
		x := x[i]             // 这里定义了局部x
		fmt.Printf("%c\n", x) // 优先用局部x
	}
	fmt.Println(x) // 使用外部的x
}
