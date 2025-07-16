package main

import "fmt"

func equalizeArray(arr []int32) int32 {
    frequencies := make(map[int32]int)

    for _, item := range arr {
        frequencies[item]++
    }

    maxFrequency := 0
    for _, count := range frequencies {
        if count > maxFrequency {
            maxFrequency = count
        }
    }

    return int32(len(arr)) - int32(maxFrequency)
}


func main() {
	fmt.Println(equalizeArray([]int32{1,2,2,3}))
}