package main

import (
	"fmt"
)

func minimumDistances(a []int32) int32 {
	var anchorNum int32 = -1
	var minDistance int32 = 2147483647
	for i := 0; i < len(a)-1; i++ {
		anchorNum = a[i]
		for j := i; j < len(a)-1; j++ {
			if anchorNum == a[j+1] {
				minDistance = min(minDistance, int32(j+1-i))
			}
		}
	}

	if minDistance == 2147483647 {
		minDistance = -1
	}
	return minDistance
}

func main() {
	fmt.Println(minimumDistances([]int32{7,1,3,4,1,7}))
}