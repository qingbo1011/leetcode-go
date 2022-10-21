package main

import "sort"

func main() {

}

func findContentChildren(g []int, s []int) int {
	ans := 0
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] { // s[j]饼干可以满足g[i]孩子，ans++，i，j都继续向后遍历
			ans++
			i++
		}
		j++
	}
	return ans
}
