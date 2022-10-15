# 438. 找到字符串中所有字母异位词

[LeetCode 438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)

## 滑动窗口

参考[潘达叔](https://leetcode.cn/u/pandaoknight/)大佬的[题解](https://leetcode.cn/problems/find-all-anagrams-in-a-string/solution/zhe-ci-xie-de-bi-jiao-shun-shou-er-qie-d-vhu9/)。

> 这里也利用了Golang中数组可以直接做比较的特性：
>
> ```go
> func main() {
> 	s := [3]int{1, 2, 3}
> 	p := [3]int{3, 2, 1}
> 	fmt.Println(s == p) // false
> }
> ```
>
> ```go
> func main() {
> 	s := [3]int{1, 2, 3}
> 	p := [3]int{1, 2, 3}
> 	fmt.Println(s == p) // true
> }
> ```

题目说了`s`和`p`仅包含小写字母，a-z对应的ASCII码范围为97-122。**所以我们维护一个长度为123的int数组：数组下标`i`对应a-z的ASCII码；数组的值`arr[i]`代表该字母出现的次数**。（其实用一个26长度的int数组也是ok的）

经过这样处理之后，比较是否出现异位词就方便多了！接下来只需要通过双指针来维护滑动窗口即可。

草稿和之前的滑动窗口题目类似，这里就不再写了。根据代码就能看懂。

代码如下：

```go
func findAnagrams(s string, p string) []int {
	var res []int
	lengthS := len(s)
	lengthP := len(p)
	if lengthS < lengthP {
		return res
	}
	var arrP [123]int
	for i := 0; i < lengthP; i++ {
		arrP[p[i]]++
	}
	// 双指针就位
	var arr [123]int
	left, right := 0, 0
	for right = 0; right < lengthP; right++ {
		arr[s[right]]++
	}
	// 维护滑动窗口
	right-- // right在上一个循环中退出循环后值为lengthP，所以这里需要right--
	for right < lengthS {
		if arr == arrP {
			res = append(res, left)
		}
		left++
		right++
		arr[s[left-1]]--
		if right < lengthS { // 防止right到最后一位时，s[right]溢出的情况
			arr[s[right]]++
		}
	}
	return res
}
```





