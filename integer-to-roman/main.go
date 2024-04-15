package main

import
	(
		"fmt"
		"math"
		"strconv"
	)

func main() {
	fmt.Println(intToRoman(12))
}

func intToRoman(num int) string {
    // map of integer to Roman numeral
    maps := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}

    // Initialize an empty string
    romanStr := ""

    // Convert the integer to a string to iterate through each digit
    numStr := strconv.Itoa(num)

    // Iterate through each digit and convert to Roman numeral
    for i := 0; i < len(numStr); i++ {
        // Calculate the position value
        pos := int(math.Pow10(len(numStr) - i - 1))
        digit := int(numStr[i] - '0')

        // Get the Roman numeral representation for the current digit
        switch digit {
        case 1, 2, 3:
            for j := 0; j < digit; j++ {
                romanStr += maps[1*pos]
            }
        case 4:
            romanStr += maps[1*pos] + maps[5*pos]
        case 5:
            romanStr += maps[5*pos]
        case 6, 7, 8:
            romanStr += maps[5*pos]
            for j := 0; j < digit-5; j++ {
                romanStr += maps[1*pos]
            }
        case 9:
            romanStr += maps[1*pos] + maps[10*pos]
        }
    }

    return romanStr
}

