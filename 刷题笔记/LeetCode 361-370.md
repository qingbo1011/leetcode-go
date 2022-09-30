# 367. 有效的完全平方数

[LeetCode 367. 有效的完全平方数](https://leetcode.cn/problems/valid-perfect-square/)

## 二分查找

写过[LeetCode 69. x 的平方根](https://leetcode.cn/problems/sqrtx/)之后，这题就显得很简单了，直接上代码吧。

代码如下：

```go
func isPerfectSquare(num int) bool {
	if num == 0 || num == 1 {
		return true
	}
	left, right := 1, num/2
	for left < right {
		middle := (left + right + 1) >> 1
		if middle*middle == num {
			return true
		} else if middle*middle < num { // 更新left
			left = middle
		} else { // 更新right
			right = middle - 1
		}
	}
	return false
}
```













