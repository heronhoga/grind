package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	energy := 100
	n := len(c)
	i := 0

	for {
		i = (i + int(k)) % n
		if c[i] == 1 {
			energy -= 3
		} else {
			energy -= 1
		}
		if i == 0 {
			break
		}
		fmt.Println("Energy: ", energy)
	}

	return int32(energy)
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3))
}