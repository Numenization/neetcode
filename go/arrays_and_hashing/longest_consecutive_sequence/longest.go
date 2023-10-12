package main

import "fmt"

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	best := 0
	for _, num := range nums {
		if m[num-1] {
			continue
		}

		sequence := 1
		temp := 1
		for m[num+temp] {
			sequence++
			temp++
		}

		if sequence > best {
			best = sequence
		}
	}

	return best
}

func main() {
	n1 := []int{100, 4, 200, 1, 3, 2}
	n2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(n1), longestConsecutive(n2))
}
