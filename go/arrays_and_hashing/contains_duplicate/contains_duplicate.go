package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	visited := make(map[int]int)

	for _, num := range nums {
		if _, ok := visited[num]; ok {
			return true
		}
		visited[num]++
	}

	return false
}

func main() {
	s1 := containsDuplicate([]int{1, 2, 3, 1})                   // true
	s2 := containsDuplicate([]int{1, 2, 3, 4})                   // false
	s3 := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) // true
	fmt.Println(s1, s2, s3)
}
