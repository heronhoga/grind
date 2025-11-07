package main

import "fmt"


func workbook(n int32, k int32, arr []int32) int32 {
	var page, specialCount int32

	page = 1
	for i := int32(0); i < n; i++ {
		problems := arr[i]
		for p := int32(1); p <= problems; p++ {
			if p == page {
				specialCount++
			}
			if p%k == 0 && p != problems {
				page++
			}
		}
		page++
	}
	return specialCount
}


func main() {
	fmt.Println("test")
	workbook(5, 3, []int32{4,2,6,1,10})
}