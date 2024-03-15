package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	
	stringInt := strconv.Itoa(x)
	stringRev := ""
	for i := len(stringInt) - 1; i >= 0; i-- {
		if stringInt[i] == '-' {
			break;
		} else {
			stringRev += string(stringInt[i])
		}
	}

	intRev, _ := strconv.Atoi(stringRev)
	if x < 0 {
		intRev = -intRev
	}

	if intRev > 2147483647 || intRev < -2147483648 {
		return 0
	}

	return intRev
}