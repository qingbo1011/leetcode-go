# 242. 有效的字母异位词

[LeetCode 242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/)

## 数组

全都是小写字母，所以用一个26长度的数组即可表示字符串（结合ASCII码，这种处理已经很常见了）。

不用多说什么，直接看代码吧。

代码如下：

```go
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
```







