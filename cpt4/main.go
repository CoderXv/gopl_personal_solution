package main

import (
	//"fmt"
	//"tutorial/cpt2/popcount"
	"fmt"
	"tutorial/cpt4/solution"
)

func main() {
	// 4.1
	//s1 := popcount.GetPc()
	//s2 := popcount.GetPc()
	//count := solution.Sha256Count(s1, s2)
	//fmt.Printf("4.1 : Sha256 Compare Same Bit Count is %d \n", count)

	// 4.2 Usage: go build main.go -method {sha256, sha384, sha512} -text ssssss
	//solution.ShaOutput()

	// 4.3 此处指针数组的用法与C语言非常相似
	array := [5]int{1, 2, 3, 4, 5}
	solution.ReverseRemaster(&array)
	fmt.Printf("4.3 Reverse: %x \n", array)

	// 4.4
	array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	res := solution.RotateOnce(&array2, 3)
	fmt.Printf("4.4 Rotate: %x \n", res)

	// 4.5
	array3 := []string{"bbbb", "bbbb", "bbbb", "bbbb"}
	res5 := solution.RmSameChar(array3)
	fmt.Print("4.5 Result: ", res5, "\n")

	// 4.6
	array5 := []rune("1 234567")
	res6 := solution.RmUniSame(array5)
	fmt.Print("4.6 RmUniSame ", res6, "\n")

	// 4.7
	array4 := []byte("我要搞Golang")
	res7 := solution.RevRemaster(array4)
	fmt.Print("4.7 Revert Result", res7, "\n")

}
