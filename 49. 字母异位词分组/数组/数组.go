package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	exist := make([]int, 0) // 存放在strs中已经放到res中的索引
	for i := 0; i < len(strs); i++ {
		if !isExist(i, exist) { // 该字符串并没有被放到res中
			exist = append(exist, i) // 先添加到exist中
			temp := []string{strs[i]}
			for j := i + 1; j < len(strs); j++ {
				if isAnagram(strs[i], strs[j]) {
					exist = append(exist, j) // 先添加到exist中
					temp = append(temp, strs[j])
				}
			}
			res = append(res, temp)
		}
	}
	return res
}

// 判断两个字符串是否为字母异位词
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

// 判断i是否在exist切片中
func isExist(i int, exist []int) bool {
	for _, ex := range exist {
		if i == ex {
			return true
		}
	}
	return false
}
