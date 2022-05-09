package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = fToC(f)
	fmt.Printf("boiling point = %g F or %g C\n", f, c)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
