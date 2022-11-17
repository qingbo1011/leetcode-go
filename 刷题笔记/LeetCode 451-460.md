# 454. 四数相加 II

[LeetCode 454. 四数相加 II](https://leetcode.cn/problems/4sum-ii/)

## HashMap

- [官方题解](https://leetcode.cn/problems/4sum-ii/solution/si-shu-xiang-jia-ii-by-leetcode-solution/)
- **[代码随想录B站视频](https://www.bilibili.com/video/BV1Md4y1Q7Yh/)**

这题看起来比较麻烦，但其实总体思路是跟[LeetCode 1. 两数之和](https://leetcode.cn/problems/two-sum/)一样的，使用HashMap来解决即可。

首先对这四个数组分一下组：`nums1[i]+nums2[j]`为一组，`nums3[k]+nums4[l]`为一组。在使用map时：key为`nums1[i]+nums2[j]`，value为这个值出现的次数。先通过一次双层循环得到map，再来一次双层循环来找到匹配到的结果。

思路很简单打一下草稿就能明白。

代码如下：

```go
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	mp := make(map[int]int, 0)
	n := len(nums1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mp[nums1[i]+nums2[j]]++
		}
	}
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			//if _, ok := mp[0-(nums3[k]+nums4[l])]; ok { // 说明nums1[i]+nums2[j]+nums3[k]+nums4[l]==0，即我们想要的结果
			//	ans = ans + mp[0-(nums3[k]+nums4[l])]
			//}
			ans = ans + mp[0-(nums3[k]+nums4[l])] // 可以直接这样写，因为go中如果mp[key]中的key不存在，取出来的value是0
		}
	}
	return ans
}
```









# 455. 分发饼干

[LeetCode 455. 分发饼干](https://leetcode.cn/problems/assign-cookies/)

## 贪心算法

这题是一道简单题，运用一下贪心算法很容易就能想出思路（排序+双指针+go）：

- 首先将`g`和`s`进行升序排序（自然顺序）
- 接下来使用`i`、`j`指针分别`g[0]`和`s[0]`。很容易理解：当`s[j]`饼干可以满足`g[i]`孩子时，`ans++`，i，j都继续向后遍历；反之当`s[j]`饼干不能满足`g[i]`孩子时，只是`j++`去寻找更大的饼干。

题目很简单直接看代码即可。

代码如下：

```go
func findContentChildren(g []int, s []int) int {
	ans := 0
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] { // s[j]饼干可以满足g[i]孩子，ans++，i，j都继续向后遍历
			ans++
			i++
			j++
		} else { // 当前饼干不能满足该孩子，j继续向后遍历找到更大的饼干
			j++
		}
	}
	return ans
}
```

> 上面是为了方便看出思路所以这么写的if-else，可以优化成：
>
> ```go
> func findContentChildren(g []int, s []int) int {
> 	ans := 0
> 	sort.Ints(g)
> 	sort.Ints(s)
> 	i, j := 0, 0
> 	for i < len(g) && j < len(s) {
> 		if g[i] <= s[j] { // s[j]饼干可以满足g[i]孩子，ans++，i，j都继续向后遍历
> 			ans++
> 			i++
> 		}
> 		j++
> 	}
> 	return ans
> }
> ```















