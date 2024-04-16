package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

}

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

	//DETERMINE SHORTEST STRING
    shortest := strs[0]
    for _, str := range strs {
        if len(str) < len(shortest) {
            shortest = str
        }
    }

    prefix := shortest
    prefixLength := len(prefix)

    for _, str := range strs {
        for prefix != str[0:prefixLength] {
            prefix = prefix[0 : prefixLength-1]
            prefixLength--

            if prefixLength < 1 {
                return ""
            }
        }
    }

    return prefix
}

