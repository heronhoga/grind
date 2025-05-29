package main

import (
	"fmt"
	"math"
)

	func squares(a int32, b int32) int32 {
		start := int32(math.Ceil(math.Sqrt(float64(a))))
		end := int32(math.Floor(math.Sqrt(float64(b))))
		if end < start {
			return 0
		}
		return end - start + 1
	}

func main() {
	fmt.Println(squares(3, 30))
}