package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // Output: "PAHNAPLSIIGYIR"
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows <= 0 || len(s) <= numRows {
		return s
	}

	var rows []string
	for i := 0; i < min(numRows, len(s)); i++ {
		rows = append(rows, "")
	}

	currentRow := 0
	direction := -1

	for i := 0; i < len(s); i++ {
		rows[currentRow] += string(s[i])
		if currentRow == 0 || currentRow == numRows-1 {
			direction = -direction
		}
		currentRow += direction
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
