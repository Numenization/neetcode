package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)
	for _, str := range strs {
		m := [26]int{}
		for _, ch := range str {
			m[ch-'a']++
		}
		anagrams[m] = append(anagrams[m], str)
	}

	result := [][]string{}
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}

func main() {
	s1, s2, s3 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}, []string{""}, []string{"a"}
	fmt.Println(groupAnagrams(s1), groupAnagrams(s2), groupAnagrams(s3))
}
