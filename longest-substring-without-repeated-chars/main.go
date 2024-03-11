package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	if s == " " {
		return 1
	}

	if s == "" {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	substringArray := make([]string, 0)
	var substringPackage string
	result := 0

	contains := func (s string, char byte) bool {
		for i := 0; i < len(s); i++ {
			if s[i] == char {
				return true
			}
		}
		return false
		
	}

	for indexStart := 0; indexStart < len(s); indexStart++ {
		for i := indexStart; i < len(s); i++ {
			if !contains(substringPackage, s[i]) {
				substringPackage += string(s[i])
			} else {
				substringArray = append(substringArray, substringPackage)
				substringPackage = ""
				break
			}
		}
	}

	for _, v := range substringArray {
		if len(v) > result {
			result = len(v)
		}
	}


	return result
}




