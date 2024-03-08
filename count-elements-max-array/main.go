package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxFrequencyElements([]int{1,2,3,4,5}))
}

func maxFrequencyElements (nums []int) int {
	min := nums[0]
	max := nums[0]

	//CREATE MAP FOR EACH ELEMENT IN NUMS
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		//LOGIC FOR MIN AND MAX
		if max < nums[i] {
            max = nums[i]
        }
        if min > nums[i] {
            min = nums[i]
        }
		//LOGIC FOR CREATING MAP
		if i == 0 {
			numsMap[nums[i]] = 1
		} else if i > 0 {
			if numsMap[nums[i]] == 0 {
				numsMap[nums[i]] = 1
			} else {
				numsMap[nums[i]]+= 1
				
			}
		}
	}

	totalFrequency := 0

	maxFreq := 0

	for _, i := range numsMap {
		if maxFreq < i {
            maxFreq = i
        }
	}

	for _, i := range numsMap {
		if i == maxFreq {
			totalFrequency += i
		}
	}

	fmt.Println(numsMap)

	return totalFrequency
}