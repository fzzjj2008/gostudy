package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰温度
	BoilingC      Celsius = 100     // 沸水温度
)

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func init() {
	fmt.Println("import tempconv")
}
