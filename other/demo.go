package main

import (
	"fmt"
	"tutorial/other/search"
	"tutorial/other/sort"
)

func callSort() {
	array := []int{5, 4, 3, 2, 1}

	// Bubble
	fmt.Println("Origin")
	fmt.Println(array)
	fmt.Println("Bubble Sort:")
	sort.BubbleSort(array)
	fmt.Println(array)

	// Insert
	fmt.Println("Reset.")
	array = sort.ResetArray()
	fmt.Println(array)
	fmt.Println("Insert Sort:")
	sort.InsertSort(array)
	fmt.Println(array)

	// Select
	fmt.Println("Reset.")
	array = sort.ResetArray()
	fmt.Println(array)
	fmt.Println("Selection Sort")
	sort.SelectionSort(array)
	fmt.Println(array)

	// Merge
	fmt.Println("Reset.")
	array = sort.ResetArray()
	fmt.Println(array)
	fmt.Println("Merge Sort")
	array = sort.MergeSort(array)
	fmt.Println(array)

	// Quick
	fmt.Println("Reset.")
	array = sort.ResetArray()
	fmt.Println(array)
	fmt.Println("Quick Sort")
	sort.QuickSort(array)
	fmt.Println(array)
}

func callSearch() {
	array := []int{1, 2, 3, 5, 6, 9}
	targetFalse := 4
	targetTrue := 5

	// Binary Search
	fmt.Printf("Binary Search, target true: %d, target false: %d \n", targetTrue, targetFalse)
	resTrue := search.BinarySearch(array, targetTrue)
	resFalse := search.BinarySearch(array, targetFalse)
	fmt.Printf("Search Result, target true: %t, target false: %t \n", resTrue, resFalse)

}

func callLinkList() {

}

func main() {
	callSort()
	callSearch()
}
