package main

import (
	"fmt"
	"math"
	"strconv"
)

func kaprekarNumbers(p int32, q int32) {
	isKaprekar := false
	for i := p; i <= q; i++ {
		if i >= 9 || i == 1 {
		mulRes := strconv.Itoa(int(math.Pow(float64(i), 2)))
		left, _ := strconv.Atoi(mulRes[:(len(mulRes)/2)])
		right, _ := strconv.Atoi(mulRes[(len(mulRes)/2):])

		if left+right == int(i) {
			isKaprekar = true
			fmt.Print(i, " ")
		}
		}
	}

	if (isKaprekar == false) {
		fmt.Println("INVALID RANGE")
	}
}

func main() {
	kaprekarNumbers(1,100)
}