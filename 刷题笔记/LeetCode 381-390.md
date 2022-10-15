# 383. 赎金信

[LeetCode 383. 赎金信](https://leetcode.cn/problems/ransom-note/)

## 数组

这种题没什么好说的了（这种思路在[LeetCode 76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)的滑动窗口题解中，只是一个check函数）。

`ransomNote` 和 `magazine` 由小写英文字母组成，那么26长度的数组即可。

代码如下：

```go
func canConstruct(ransomNote string, magazine string) bool {
	ransomArr, magazineArr := [26]int{}, [26]int{}
	for i := 0; i < len(ransomNote); i++ {
		ransomArr[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		magazineArr[magazine[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if ransomArr[i] > magazineArr[i] {
			return false
		}
	}
	return true
}
```



















