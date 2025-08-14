package main

import (
	"fmt"
)

func checkSimilarity(binaryArr1 string, binaryArr2 string) int {
	score := 0
	for i := 0; i < len(binaryArr1); i++ {
		if binaryArr1[i] == '1' || binaryArr2[i] == '1' {
			score++;
		}
	}

	fmt.Println("score temp : ", score)

	return score
}
func acmTeam(topic []string) []int32 {
	team := 0
	topicTotal := 0
    for i := 0; i < len(topic) - 1; i++ {
		for j := i+1; j < len(topic); j++ {
			score := checkSimilarity(topic[i], topic[j])
			if score > topicTotal {
				topicTotal = score
				team = 1
			} else if score == topicTotal {
				team++
			}
		}
	}

	return []int32{int32(topicTotal), int32(team)}
}

func main() {
	fmt.Println(acmTeam([]string{"10101", "11100", "11010", "00101"}))
}