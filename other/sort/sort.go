package sort

func ResetArray() []int {
	return []int{5, 4, 3, 2, 1}
}

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array) - i - 1; j++ {
			if array[j] > array[j + 1] {
				array[j], array[j + 1] = array[j + 1], array[j]
			}
		}
	}
}

func InsertSort(array []int) {
	for i := 0; i < len(array); i += 1 {
		value := array[i]
		// [0:i-1]的数组是稳定有序的
		j := i - 1
		for ; j >= 0; j -= 1 {
			if array[j] > value {
				// 右移
				array[j + 1] = array[j]
			} else {
				break
			}
		}
		array[j + 1] = value
	}
}

func SelectionSort(array []int) {
	// [0: i-1] 都是有序的
	for i := 0; i < len(array); i++ {
		j := i
		min := array[j]
		minIndex := j
		for ; j < len(array); j ++ {
			if min > array[j] {
				min = array[j]
				minIndex = j
			}
		}
		array[minIndex], array[i] = array[i], array[minIndex]
	}
}

func MergeSort(array []int) []int {
	return mergeSort(array)
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := (len(array)) / 2
	arrayA := mergeSort(array[0:mid])
	arrayB := mergeSort(array[mid:])
	result := merge(arrayA, arrayB)
	return result
}

func merge(arrayA []int, arrayB []int) []int {

	if len(arrayA) == 0 {
		return arrayB
	}

	if len(arrayB) == 0 {
		return arrayA
	}

	result := make([]int, 0)

	a, b := 0, 0
	lenA, lenB := len(arrayA), len(arrayB)

	for ; a < lenA && b < lenB; {
		if arrayA [a] < arrayB[b] {
			result = append(result, arrayA[a])
			a++
			continue
		} else {
			result = append(result, arrayB[b])
			b++
		}
	}

	if a == len(arrayA) {
		result = append(result, arrayB[b:]...)
	} else {
		result = append(result, arrayA[a:]...)
	}

	return result
}

func QuickSort(array []int) {
	quickSort(array, 0, len(array) - 1)

}

func quickSort(array []int, left int, right int) {
	if left >= right {
		return
	}

	pivot := parition(array, left, right)
	quickSort(array, left, pivot - 1)
	quickSort(array, pivot + 1, right)
}

func parition(array []int, start int, end int) int {
	pivotValue := array[end]
	pivot := start
	// 交换位置，右边的是大于pivot的数字，左边是小于pivot的数字
	for j := start; j < end; j++ {
		if array[j] < pivotValue {
			array[pivot], array[j] = array[j], array[pivot]
			pivot++
		}
	}
	array[pivot], array[end] = array[end], array[pivot]
	return pivot
}


