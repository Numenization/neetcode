package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	r := []int{}
	for i := 0; i < k; i++ {
		top := 0
		key := 0
		for j, v := range freq {
			if v > top {
				top = v
				key = j
			}
		}
		r = append(r, key)
		freq[key] = 0
	}

	return r
}

func main() {
	n1, k1 := []int{1, 1, 1, 2, 2, 3}, 2
	n2, k2 := []int{1}, 1

	fmt.Println(topKFrequent(n1, k1), topKFrequent(n2, k2))
}
