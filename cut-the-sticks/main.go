package main

import "fmt"

//sort
func partition(arr []int32, low, high int) ([]int32, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int32, low, high int) []int32 {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int32) []int32 {
	return quickSort(arr, 0, len(arr)-1)
}
//end sort

func cutTheSticks(arr []int32) []int32 {
	sortedArr := quickSortStart(arr)
	var resArr []int32

	for len(sortedArr) != 0 {
		resArr = append(resArr, int32(len(sortedArr)))
		currentVal := sortedArr[0]
		for i := 0; i < len(sortedArr); i++ {
			sortedArr[i] -= currentVal
		}
		cursor := 0
		for cursor < len(sortedArr) && sortedArr[cursor] == 0 {
			cursor++
		}
		sortedArr = append([]int32{}, sortedArr[cursor:]...)
	}
	return resArr
}



func main() {
	fmt.Println(cutTheSticks([]int32{5, 6, 7, 2, 1, 0}))
}