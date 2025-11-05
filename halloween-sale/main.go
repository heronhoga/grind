package main

import "fmt"

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	basePrice := p
	wallet := s
	count := 0

	//if wallet is less than base price
	if basePrice > wallet {
		return 0
	}

	for wallet >= basePrice {
			wallet -= basePrice
			count++
			basePrice -= d
			if basePrice < m {
				basePrice = m
			}
	}

	return int32(count)
}

func main() {
	fmt.Println(howManyGames(20, 3,6,80))
}