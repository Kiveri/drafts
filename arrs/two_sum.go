package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target

func twoSum(nums []int, target int) []int {
	sum := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sum[0] = i
				sum[1] = j
			}
		}
	}
	return sum
}

func main() {
	nums := []int{1, 3, 9, 4}
	target := 5
	fmt.Println(twoSum(nums, target))
}
