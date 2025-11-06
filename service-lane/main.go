package main

import "fmt"

func serviceLane(n int32, t int32, width []int32, cases [][]int32) []int32 {
	result := []int32{}
	for i := 0; i < len(cases); i++ {
		start := cases[i][0]
		end := cases[i][1]

		minVal := width[start]
		for i := start; i <= end; i++ {
			if width[i] < minVal {
				minVal = width[i]
			}
		}

		result = append(result, minVal)

	}

	return result
}

func main() {
	fmt.Println(	serviceLane(4, 2, []int32{2, 3, 1, 2, 3, 2, 3, 3}, [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}))
}