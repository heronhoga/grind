package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    var characters = [8]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    var result []string

    var generateCombinations func(index int, currentCombination string)

    generateCombinations = func(index int, currentCombination string) {
        if index == len(digits) {
            result = append(result, currentCombination)
            return
        }

        digitIndex := int(digits[index] - '0') - 2
        characterSet := characters[digitIndex]

        for _, char := range characterSet {
            generateCombinations(index+1, currentCombination+string(char))
        }
    }

    generateCombinations(0, "")

    return result
}
