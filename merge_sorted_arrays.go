package main

import "fmt"

func MergeSortedArrays(arr1, arr2 []int) []int {
	mergedArray := make([]int, 0)
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArray = append(mergedArray, arr1[i])
			i++
		} else {
			mergedArray = append(mergedArray, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		mergedArray = append(mergedArray, arr1[i])
		i++
	}

	for j < len(arr2) {
		mergedArray = append(mergedArray, arr2[j])
		j++
	}

	return mergedArray
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	array := []int{6, 3, 8, 2, 9, 1}
	bubbleSort(array)
	fmt.Println("Sorted array:", array)

	array1 := []int{1, 3, 5, 7}
	array2 := []int{2, 4, 6, 8, 9}
	merged := MergeSortedArrays(array1, array2)

	fmt.Println("Merged sorted arrays:", merged)
}
