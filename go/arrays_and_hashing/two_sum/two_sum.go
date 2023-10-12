package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok && i != j {
			return []int{j, i}
		}
		m[v] = i
	}

	return []int{}
}

func main() {
	n1, t1 := []int{2, 7, 11, 15}, 9
	n2, t2 := []int{3, 2, 4}, 6
	n3, t3 := []int{3, 3}, 6

	r1, r2, r3 := twoSum(n1, t1), twoSum(n2, t2), twoSum(n3, t3)
	fmt.Println(r1, r2, r3) // [0,1] [1,2] [0,1]
}
