package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("1a"))
}

func myAtoi(s string) int {
    cleanStringInt := strings.TrimSpace(s)
    // Handle leading sign
    var sign int = 1
    if len(cleanStringInt) > 0 && (cleanStringInt[0] == '-' || cleanStringInt[0] == '+') {
        if cleanStringInt[0] == '-' {
            sign = -1
        }
        cleanStringInt = cleanStringInt[1:] // Remove the sign from the string
    }

    // Process remaining characters
    newString := ""
    for _, char := range cleanStringInt {
        if char >= '0' && char <= '9' {
            newString += string(char)
        } else {
            break
        }
    }

    // Convert string to int
    intResult, _ := strconv.Atoi(newString)

    // Apply sign
    intResult *= sign

    // Clamp the result if it's outside the range of int32
    if intResult < -(1<<31) {
        return -(1 << 31)
    }
    if intResult > (1<<31) - 1 {
        return (1<<31) - 1
    }

    return intResult
}