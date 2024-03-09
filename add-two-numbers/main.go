package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := &ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{9, nil}}}}

	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Println("nil")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	reverseString := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	var num1Str string
	for l1 != nil {
		num1Str += fmt.Sprintf("%d", l1.Val)
		fmt.Println(num1Str)
		l1 = l1.Next
	}

	var num2Str string
	for l2 != nil {
		num2Str += fmt.Sprintf("%d", l2.Val)
		fmt.Println(num2Str)
		l2 = l2.Next
	}

	num1Str = reverseString(num1Str)
	num2Str = reverseString(num2Str)
	num1 := new(big.Int)
	num2 := new(big.Int)
	num1.SetString(num1Str, 10)
	num2.SetString(num2Str, 10)

	sum := new(big.Int).Add(num1, num2)
	sumString := sum.String()

	var result *ListNode

	for i := 0; i < len(sumString); i++ {
		digit, _ := strconv.Atoi(string(sumString[i]))
		newNode := &ListNode{
			Val:  digit,
			Next: result,
		}
		result = newNode
	}

	return result
}
