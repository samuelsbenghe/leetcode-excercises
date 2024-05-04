// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Explain the rules
	fmt.Println("Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.")
	fmt.Println("You may assume that each input would have exactly one solution, and you may not use the same element twice.")
	fmt.Println("You can return the answer in any order.")
	// Test cases taken from leetcode problem
	fmt.Printf("Given nums = [2,7,11,15], target = 9, the two indices are: %v\n", twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Printf("Given nums = [3,2,4], target = 6, the two indices are: %v\n", twoSum([]int{3, 2, 4}, 6))          // [1, 2]
	fmt.Printf("Given nums = [3,3], target = 6, the two indices are: %v\n", twoSum([]int{3, 3}, 6))               // [0, 1]
	// Ask the user for input until they enter an empty string. This is to allow the user to test their own arrays.
	for {
		fmt.Println("Enter an array of integers separated by commas (e.g. 1,2,3,4):")
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		fmt.Println("Enter a target integer:")
		var target int
		fmt.Scanln(&target)
		fmt.Printf("The two indices are: %v\n", twoSum(parseInput(input), target))
	}
}

func parseInput(input string) []int {
	var nums []int
	for _, numStr := range strings.Split(input, ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input, please enter integers separated by commas.")
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			} else {
				if nums[i]+nums[j] == target {
					result := []int{i, j}
					return result
				}
			}
		}
	}
	return nil
}
