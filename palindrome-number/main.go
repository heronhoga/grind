package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
    stringInt := strconv.Itoa(x)

	length := len(stringInt)

	for i := 0; i < length / 2; i++ {
		if stringInt[i] != stringInt[length - 1 - i] {
			return false
		}
	}

	return true
}