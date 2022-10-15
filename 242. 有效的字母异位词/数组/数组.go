package main

func main() {

}

func isAnagram(s string, t string) bool {
	lengthS, lengthT := len(s), len(t)
	if lengthS != lengthT {
		return false
	}
	sArr, tArr := [26]int{}, [26]int{}
	for i := 0; i < lengthS; i++ {
		sArr[s[i]-'a']++
		tArr[t[i]-'a']++
	}
	return sArr == tArr // go中，数组是可以直接比较的
}
