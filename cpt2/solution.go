package main

import (
	"fmt"
	"tutorial/cpt2/tempconv"
	"tutorial/cpt2/popcount"
)

// 2.1
func solution1() {
	// 2.1
	fmt.Println("C to K at Zero :", tempconv.Freezing.String(), "equals", tempconv.CToK(tempconv.Freezing).String())
}

//2.2
func solution2() {
	res := popcount.PopCount(111)
	fmt.Printf("PopCount: %d \n", res)

	res = popcount.PopCountLoop(111)
	fmt.Printf("PopCountLoop: %d \n", res)

	res = popcount.PopCount2(111)
	fmt.Printf("PopCount2: %d \n", res)

	res = popcount.PopCount3(111)
	fmt.Printf("PopCount3: %d \n", res)
}
