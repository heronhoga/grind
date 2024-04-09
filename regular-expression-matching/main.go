package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a*"))
}

func isMatch(s string, p string) bool {

	//return true if only s is empty
    if p == "" {
        return s == ""
    }

    firstMatch := s != "" && (p[0] == s[0] || p[0] == '.')

    if len(p) >= 2 && p[1] == '*' {
        return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
    } else {
        return firstMatch && isMatch(s[1:], p[1:])
    }
}
