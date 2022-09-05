# 55. 跳跃游戏

[LeetCode 55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)

## 贪心算法

这里，我们用贪心的思想，求出前**n-1个数**（**注意是前n-1个数**，因为最后一个数是多少并不重要，只要能到达最后一个数所在的下标即可。n = nums.length）。能到达的下标的最大值farthest是否>=n-1（**n-1即为nums最后一个数所在下标**）。

前n-1个数在代码中体现在：`for i := 0; i < n-1; i++`

```go
for i := 0; i < n-1; i++ {
...
}
```

分析：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220502150507.png)

其中注意中间卡住的判断：**如果最大距离（所在下标）都小于或等于i了，说明i之前的数字无论如何也无法跳跃到i之后（中途遇到0卡住了）**，代码体现在：

```go
if farthest <= i {
   return false
}
```

代码如下：

```go
func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i { // 如果最大距离（所在下标）都小于或等于i了，说明i之前的数字无论如何也无法跳跃到i之后（中途遇到0卡住了）
			return false
		}
	}
	return farthest >= n-1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
```







