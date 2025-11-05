package main

import "fmt"

func beautifulTriplets(d int32, arr []int32) int32 {
	set := make(map[int32]bool)
	count := 0
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
	}

	for i := 0; i < len(arr); i++ {
		if set[arr[i] + d] && set[arr[i] + (2*d)] {
			count++
		}
	}

	return int32(count)
}

func main() {
	fmt.Println(beautifulTriplets(3, []int32{1,2,4,5,7,8,10}))
}
