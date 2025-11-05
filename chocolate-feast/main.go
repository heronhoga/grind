package main

import "fmt"

func chocolateFeast(n int32, c int32, m int32) int32 {
	feast := n / c
	wrappers := n / c
	for wrappers >= m {
		wrappers = wrappers - m + 1
		feast++
	}

	return feast
}

func main() {
	fmt.Println(chocolateFeast(15, 3 ,2))
}