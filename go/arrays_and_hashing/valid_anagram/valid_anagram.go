package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	visited := make(map[byte]int)

	for i := range s {
		visited[s[i]]++
		visited[t[i]]--
	}

	for _, v := range visited {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	s1, t1 := "anagram", "nagaram"
	s2, t2 := "aacc", "ccac"
	r1, r2 := isAnagram(s1, t1), isAnagram(s2, t2)
	fmt.Println(r1, r2) // true false
}
