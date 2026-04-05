package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	storedMap := make(map[string][]string)

	for _, str := range strs {
		bytes := []byte(str)
		slices.Sort(bytes)
		storedMap[string(bytes)] = append(storedMap[string(bytes)], str)
	}

	res := make([][]string, 0, len(storedMap))
	for _, value := range storedMap {
		res = append(res, value)
	}

	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "nat", "bat"}))
}
