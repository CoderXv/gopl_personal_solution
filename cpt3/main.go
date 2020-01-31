package main

import (
	"fmt"
	"tutorial/cpt3/solution"
)

func main() {

	// test
	// print("hellow world")

	// 3.1 go run main.go > output.xml
	// solution.Solution1()

	// 3.2 omit

	// 3.3 go run main.go > output.xml
	// solution.Solution3()

	// 3.4 omit

	// 3.5
	// solution.Mandelbrot()

	// 3.6 - 3.9 omit

	// 3.10
	//res10 := solution.Comma("1234567890")
	//fmt.Println(res)

	// 3.11
	//res11 := solution.CommaFloatFlag("+123456.7890")
	//fmt.Println(res11)

	// 3.12
	res12 := solution.Isomerism("adc", "dac")
	fmt.Println(res12)
}
