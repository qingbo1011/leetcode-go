package main

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	pArr := [26]int{}
	for _, bytes := range p {
		pArr[bytes-'a']++
	}
	for left := 0; left < len(s)-len(p)+1; left++ {
		right := left + len(p) - 1
		sArr := [26]int{}
		for i := left; i <= right; i++ {
			sArr[s[i]-'a']++
		}
		if sArr == pArr {
			res = append(res, left)
		}
	}

	return res
}

func main() {

}
