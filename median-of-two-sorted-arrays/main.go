package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arrMerged := Merge(nums1, nums2)
	fmt.Println("array digabung")
	fmt.Println(arrMerged)
	arrSorted := MergeSort(arrMerged)
	fmt.Println("array disort")
	fmt.Println(arrSorted)
	if len(arrSorted)%2 == 0 {
		return float64(arrSorted[len(arrSorted)/2-1]+arrSorted[len(arrSorted)/2]) / 2
	}

	return float64(arrSorted[len(arrSorted)/2])
}

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	
	middle := len(array) / 2
	left := MergeSort(array[:middle])
	fmt.Println(left)
	right := MergeSort(array[middle:])
	fmt.Println(right)

	return Merge(left, right)
}

func Merge(arrA []int, arrB []int) []int {
	arrResult := make([]int, 0)
	arrayA := arrA
	arrayB := arrB

	for len(arrayA) > 0 && len(arrayB) > 0 {
		if arrayA[0] > arrayB[0] {
			arrResult = append(arrResult, arrayB[0])
			arrayB = arrayB[1:]
		} else {
			arrResult = append(arrResult, arrayA[0])
			arrayA = arrayA[1:]
		}
	}

	for len(arrayA) > 0 {
		arrResult = append(arrResult, arrayA[0])
		arrayA = arrayA[1:]
	}

	for len(arrayB) > 0 {
		arrResult = append(arrResult, arrayB[0])
		arrayB = arrayB[1:]
	}

	return arrResult
}