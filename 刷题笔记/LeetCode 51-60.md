# 53. 最大子数组和

[LeetCode 53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)

## 贪心算法

动态规划的解法等到后面写动态规划笔记时再补充。这里先用贪心算法解决。

这题需要用贪心思想简单分析，就可以得出时间复杂度为O(n)的算法，而不是完全暴力的O(n^2)。

贪心思想：在遍历过程中，**如果前面几个数加起来的和sum<0了，那下一个数当然不愿意加上前面的和，所以就舍去前面的子数组，重新开始计算连续子数组和**。

我们用max来存储最大值，sum来存储遍历过程中子数组的和。在遍历过程中：

- 如果`sum<0`，那么就舍去前面的数，从i开始继续向下遍历（`sum=nums[i]`）
- 如果`sum>=0`，继续向下遍历即可
- 注意更新`max`

代码如下：

```go
func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, num := range nums {
		if sum < 0 {
			sum = num
		} else {
			sum = sum + num
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
```



# 54. 螺旋矩阵

[LeetCode 54. 螺旋矩阵](https://leetcode.cn/problems/spiral-matrix/)

## 模拟

这题在模拟上基本上参考了在写[LeetCode 59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/)时模拟二的题解。

具体思路：

- m行n列：`m, n := len(matrix), len(matrix[0])`。初始赋值，`left, right := 0, n-1`，   `top, bottom := 0, m-1`
- 第1条边，从左到右，遍历过程为` for i := left; i <= right; i++`，赋值坐标为`[top][i]`，遍历结束后需要`top++`（这样就可以准备遍历第2条边了，画一下图就很好理解）
- 第2条边，从上到下，遍历过程为` for i := top; i <= bottom; i++`，赋值坐标为`[i][right]`，遍历结束后需要`right--`
- 第3条边，从右到左，遍历过程为`for i := right; i >= left; i--`，赋值坐标为`[bottom][i]`，遍历结束后需要`bottom--`
- 第4条边，从下到上，遍历过程为`for i := bottom; i >= top; i--`，赋值坐标为`[i][left]`，遍历结束后需要`left++`

遍历的次数为`m*n`，所以最外层有一个`for count < m*n`。需要注意的是，因为给定的不是正方形，而是m行n列的矩形，所以在每一条边遍历完成后都要判断一下是否遍历结束了。即：

```go
if count >= m*n {
   break
}
```

代码如下：

```go
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	m, n := len(matrix), len(matrix[0]) // m行n列
	// 初始化一些值
	left, right := 0, n-1
	top, bottom := 0, m-1
	count := 0
	// 开始遍历
	for count < m*n { // 一共遍历m*n次
		for i := left; i <= right; i++ { // 第1条边，从左到右
			res = append(res, matrix[top][i])
			count++
		}
		if count >= m*n {
			break
		}
		top++
		for i := top; i <= bottom; i++ { // 第2条边，从上到下
			res = append(res, matrix[i][right])
			count++
		}
		if count >= m*n {
			break
		}
		right--
		for i := right; i >= left; i-- { // 第3条边，从右到左
			res = append(res, matrix[bottom][i])
			count++
		}
		if count >= m*n {
			break
		}
		bottom--
		for i := bottom; i >= top; i-- { // 第4条边，从下到上
			res = append(res, matrix[i][left])
			count++
		}
		left++
	}
	return res
}
```



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



# 59. 螺旋矩阵 II

[LeetCode 59. 螺旋矩阵 II](https://leetcode.cn/problems/spiral-matrix-ii/)

## 模拟：方法1

参考代码随想录：

- [螺旋矩阵 II](https://programmercarl.com/0059.%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5II.html#_59-%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5ii)
- **[B站视频](https://www.bilibili.com/video/BV1SL4y1N7mV/?)**

按照B站视频的思路，采取左闭右开的区间，转`n/2`圈，每一圈遍历4条边。具体草稿就不贴了，详细的分析参考视频。

代码如下：

```go
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	// 初始化res
	var res [][]int
	for i := 0; i < n; i++ { // 先把正方形框架定下来（因为是切片，不然后面直接赋值会out of index）
		res = append(res, make([]int, n))
	}
	// 定义一些变量
	count := 1                 // 赋值给螺旋矩阵的具体值
	i, j := 0, 0               // 遍历过程中具体的坐标
	startX, startY := 0, 0     // 每一圈遍历时，开始的坐标
	offset := 1                // 在每一圈的赋值中控制边界
	for k := 0; k < n/2; k++ { // 赋值需要转n/2圈
		for j = startY; j < n-offset; j++ { // 向右遍历，第1条边
			res[startX][j] = count
			count++
		}
		for i = startX; i < n-offset; i++ { // 向下遍历，第2条边
			res[i][j] = count // 第1条边遍历结束后，j已经在第2条边的位置了
			count++
		}
		for ; j > startY; j-- { // 向左遍历，第3条边（不指定j的值是因为，第1条边遍历结束后，j已经在指定位置了）
			res[i][j] = count
			count++
		}
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}
		// 4条边遍历完成后，一圈就转完了，更新相关数据
		startX++
		startY++
		offset++
		j++
	}
	if n%2 != 0 { // 如果n是奇数，那么转n/2圈后最中间的值还是空的，此时需要赋值
		i++ // 别忘了i++（因为每条边的遍历都是左闭右开嘛）
		res[i][j] = count
	}
	return res
}
```

## 模拟：方法2

但是我在代码随想录的文字版中Golang版本看到了不一样的解法，主要思路是使用`left`，`right`来控制横向坐标，使用`top`，`bottom`来控制纵向坐标。

具体思路：

- 初始赋值，`left, right := 0, n-1`，`top, bottom := 0, n-1`
- 第1条边，从左到右，遍历过程为` for i := left; i <= right; i++`，赋值坐标为`[top][i]`，遍历结束后需要`top++`（这样就可以准备遍历第2条边了，画一下图就很好理解）
- 第2条边，从上到下，遍历过程为` for i := top; i <= bottom; i++`，赋值坐标为`[i][right]`，遍历结束后需要`right--`
- 第3条边，从右到左，遍历过程为`for i := right; i >= left; i--`，赋值坐标为`[bottom][i]`，遍历结束后需要`bottom--`
- 第4条边，从下到上，遍历过程为`for i := bottom; i >= top; i--`，赋值坐标为`[i][left]`，遍历结束后需要`left++`

至于上述遍历的次数也很好把握，初始赋值`count := 1`，count最终肯定为`n*n`，所以只需要在最外层   `for count <= n*n`即可。

代码如下：

```go
func generateMatrix(n int) [][]int {
	// 初始化res
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	// 定义一些初始值
	left, right := 0, n-1
	top, bottom := 0, n-1
	count := 1
	for count <= n*n {
		for i := left; i <= right; i++ { // 第1条边，从左到右
			res[top][i] = count
			count++
		}
		top++                            // 准备遍历下一条边
		for i := top; i <= bottom; i++ { // 第2条边，从上到下
			res[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- { // 第3条边，从右到左
			res[bottom][i] = count
			count++
		}
		bottom--
		for i := bottom; i >= top; i-- { // 第4条边，从下到上
			res[i][left] = count
			count++
		}
		left++
	}
	return res
}
```



























