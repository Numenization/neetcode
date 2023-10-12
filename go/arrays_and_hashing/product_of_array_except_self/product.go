package main

import "fmt"

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	pre := 1
	for i, num := range nums {
		res[i] = pre
		pre *= num
	}

	suf := 1
	for j := len(res) - 1; j >= 0; j-- {
		res[j] *= suf
		suf *= nums[j]
	}

	return res
}

func main() {
	n1, n2 := []int{1, 2, 3, 4}, []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(n1), productExceptSelf(n2))
	// expecting:
	// [24,12,8,6] [0,0,9,0,0]
}
