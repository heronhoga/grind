package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	currentPos := 0
	counter := 0
	for currentPos != len(c)-1 {
		if currentPos+2 <= len(c)-1 && c[currentPos+2] != 1 {
			currentPos += 2
			counter++
		} else if currentPos+1 <= len(c)-1 && c[currentPos+1] != 1 {
			currentPos += 1
			counter++
		}
	}
	return int32(counter)

}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0,0,0,0,1,0}))
}