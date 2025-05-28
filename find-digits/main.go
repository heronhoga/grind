package main

import (
	"fmt"
	"strconv"
)

func findDigits(n int32) int32 {
    digitString := strconv.Itoa(int(n))
	var result int32 = 0
	for i := 0; i < len(digitString); i++ {
		digitInt , _ := strconv.Atoi(string(digitString[i]))
		if int32(digitInt) == 0 {
			continue
		}
		if n % int32(digitInt) == 0 {
			result++
		}
	}
	
	return result
}

func main() {
	fmt.Println(findDigits(1012))
}