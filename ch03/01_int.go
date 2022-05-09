package main

import (
	"fmt"
	"unsafe"
)

func test01() {
	var apples int32 = 1
	var oranges int64 = 2
	// sum := apples + oranges // 类型不同报错
	sum := int(apples) + int(oranges)
	fmt.Println(sum)
}

func test02() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // 22269 国 '国'
	fmt.Printf("%d %[1]c %[1]q\n", newline) // 10 \n '\n'
	fmt.Println(unsafe.Sizeof(unicode))     // 4
}

func main() {
	test01()
	test02()
}
