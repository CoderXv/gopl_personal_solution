package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 1.1
func solution1() {
	fmt.Println("Solution 1.1: Print os.Args[0]: " + os.Args[0])
}

// 1.2
func solution2() {
	fmt.Println("Solution 1.2: Print index as well as os.Args.")
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("Index: " + strconv.Itoa(i) + ", os.Args: " + os.Args[i])
	}
}

// 1.3
func solution3() {
	//var withoutJoin string
	var args []string

	// init string array
	for i := 0; i < 1000; i++ {
		args = append(args, "x")
	}

	var args2 []string
	for i := 0; i < 1000; i++ {
		args2 = append(args2, "x")
	}

	// join
	startTime := time.Now()
	strings.Join(args, " ")
	endTime := time.Now()
	joinConsume := endTime.Sub(startTime)

	fmt.Print("Join Comsume: ")
	fmt.Println(joinConsume)

	// concat
	startTime = time.Now()
	var s, sep string
	for _, arg := range args2[1:] {
		s += sep + arg
		sep = " "
	}

	endTime = time.Now()
	concatConsume := endTime.Sub(startTime)

	fmt.Print("Low Efficient Concat Comsume: ")
	fmt.Println(concatConsume)

}

// main
func main() {

	// 1.1 go run main.go xxx xxx
	// solution1()

	// 1.2 go run main.go xxx xxx
	// solution2()

	//1.3 go run main.go xxx xxx
	// solution3()

	// 1.4 go run main.go solution.go  /Users/user/go/src/tutorial/cpt1/config/test
	// solution4()

	// 1.5 & 1.6 go build main.go solution.go && ./main >out.gif
	//solution5()

	// 1.7 1.8 1.9  go run main.go solution.go www.gopl.io
	//solution6()

	// 1.10 1.11 go run main.go solution.go http://www.gopl.io http://www.godoc.org
	solution7()
}

