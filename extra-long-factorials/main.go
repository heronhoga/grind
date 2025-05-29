package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	result := big.NewInt(1)
	for i := int64(2); i <= int64(n); i++ {
		result.Mul(result, big.NewInt(i))
	}
	fmt.Println(result)
}

func main() {
	extraLongFactorials(1)
}