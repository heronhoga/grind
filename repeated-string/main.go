package main

import (
	"fmt"
	"math"
)

func repeatedString(s string, n int64) int64 {
	//define result variable
	calcResult := 0
	//calculate 'a' for a string only
	aInString := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			aInString++
		}
	}

	//multiply
	nInt := int(n)
	divisionResult := nInt/len(s)
	aMultiplier := math.Floor(float64(divisionResult))

	//add result with multiplier
	calcResult += aInString * int(aMultiplier)

	//check remainder
	remainder := nInt % len(s)

	if remainder != 0 {
		for i := 0; i < remainder; i++ {
			if s[i] == 'a' {
				calcResult++
			}
		}
	}

	return int64(calcResult)

}

func main() {
	fmt.Println(repeatedString("a", 1000000000000))
}