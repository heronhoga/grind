package main

import "fmt"

func appendAndDelete(s string, t string, k int32) string {
	commonLength := 0
	minLength := len(s)
	if len(t) < minLength {
		minLength = len(t)
	}

	for i := 0; i < minLength; i++ {
		if s[i] == t[i] {
			commonLength++
		} else {
			break
		}
	}

	operationsNeeded := (len(s) - commonLength) + (len(t) - commonLength)

	if int(k) == len(s)+len(t) {
		return "Yes"
	} else if int(k) >= operationsNeeded && (int(k)-operationsNeeded)%2 == 0 {
		return "Yes"
	} else if int(k) >= len(s)+len(t) {
		return "Yes"
	} else {
		return "No"
	}
}


func main() {
	fmt.Println(appendAndDelete("hackerhappy", "hackerrank", 9))
}