package main

import "fmt"

func nonDivisibleSubset(k int32, s []int32) int32 {
    count := make([]int32, k)

    for _, num := range s {
        remainder := num % k
        count[remainder]++
    }

    var ans int32 = 0

    if count[0] > 0 {
        ans++
    }

    if k%2 == 0 && count[k/2] > 0 {
        ans++
    }

    for i := int32(1); i <= k/2; i++ {
        if i != k-i {
            if count[i] > count[k-i] {
                ans += count[i]
            } else {
                ans += count[k-i]
            }
        }
    }

    return ans
}


func main() {
	fmt.Println(nonDivisibleSubset(4, []int32{19, 10, 12, 10, 24, 25, 22}))
}