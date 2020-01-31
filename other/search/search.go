package search

import (
	"tutorial/other/sort"
)

func BinarySearch(array []int, target int) bool {
	sort.QuickSort(array)
	return binarySearch(array, target, 0, len(array) - 1)
}

func binarySearch(array []int, target int, left int, right int) bool {
	if left > right {
		return array[left] == target
	}

	if target == array[left] {
		return true
	}

	mid := left + (right - left) >> 1

	if target == array[mid] {
		return true
	} else if target > array[mid] {
		return binarySearch(array, target, mid + 1, right)
	} else {
		return binarySearch(array, target, left, mid - 1)
	}

}