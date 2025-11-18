package main

import (
	"fmt"
	"math"
)

func flatlandSpaceStations(n int32, c []int32) int32 {
	maxVal := -2147483648
	for i := 0; i < int(n); i++ {
		var minVal int32 = 2147483647

		for j := 0; j < len(c); j++ {
			if int32(i) == c[j] {
				minVal = 0
			}

			calc := c[j] - int32(i)
			calc = int32(math.Abs(float64(calc)))
			if calc < minVal {
				minVal = calc
			}
		}

		if minVal > int32(maxVal) {
			maxVal = int(minVal)
		}
	}

	return int32(maxVal)
}

func main() {
	fmt.Println(flatlandSpaceStations(5, []int32{0,4} ))
}