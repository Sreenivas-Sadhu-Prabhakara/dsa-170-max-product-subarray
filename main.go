package main

import "fmt"

// MaxProduct returns the largest product over all non-empty contiguous
// subarrays. Kadane's shape, but tracking BOTH the max and min product
// ending here — a negative value can flip the minimum into the maximum.
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxHere, minHere, best := nums[0], nums[0], nums[0]
	for _, v := range nums[1:] {
		candidates := [3]int{v, v * maxHere, v * minHere}
		maxHere, minHere = candidates[0], candidates[0]
		for _, c := range candidates[1:] {
			if c > maxHere {
				maxHere = c
			}
			if c < minHere {
				minHere = c
			}
		}
		if maxHere > best {
			best = maxHere
		}
	}
	return best
}

func main() {
	fmt.Println(MaxProduct([]int{2, 3, -2, 4}))
	fmt.Println(MaxProduct([]int{-2, 3, -4}))
}
