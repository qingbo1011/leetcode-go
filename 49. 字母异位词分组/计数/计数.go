package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		array := [26]int{} // 都是小写，所以可以用长度26的数组
		for _, i := range str {
			array[i-'a']++
		}
		mp[array] = append(mp[array], str)
	}

	res := make([][]string, 0, len(mp))
	for _, strings := range mp {
		res = append(res, strings)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "nat", "bat"}))
}
