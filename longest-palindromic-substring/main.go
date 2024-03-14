package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
    var substring string = ""

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if isPalindrome(s[i:j]) && len(s[i:j]) > len(substring) {
				substring = s[i:j]
			}
		}
	}

	return substring
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

//MORE EFFICIENT SOLUTION

// func longestPalindrome(s string) string {
//     if s == "" {
//         return ""
//     }

//     start := 0
//     end := 0

//     for i :=0; i < len(s); i++ {
//         len1 := expandAroundCenter(s, i, i)
//         len2 := expandAroundCenter(s, i, i+1)
//         maxLen := max(len1, len2)

//         if maxLen > end - start {
//             start = i - (maxLen - 1) / 2
//             end = i + maxLen / 2
//         }
//     }
//     runes := []rune(s)
//     substring := string(runes[start:end+1])
//     return substring
// }

// func expandAroundCenter(s string, left int, right int) int {
//     L := left
//     R := right
//     for L >= 0 && R < len(s) && s[L] == s[R] {
//         L -= 1
//         R += 1
//     }
//     return R - L - 1
// }