package main

import (
	"fmt"

	"github.com/fzzjj2008/gostudy/ch02/tempconv"
)

func main() {
	const temperate = 100.0
	f := tempconv.CToF(temperate)
	c := tempconv.FToC(temperate)
	fmt.Printf("%.1f C = %.1f F\n", temperate, f)
	fmt.Printf("%.1f F = %.1f C\n", temperate, c)
}
