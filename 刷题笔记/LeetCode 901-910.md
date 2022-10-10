# 904. 水果成篮

[LeetCode 904. 水果成篮](https://leetcode.cn/problems/fruit-into-baskets/)

## 滑动窗口

理解完题意，就可以使用滑动窗口的思想来解题了。

因为只有两个篮子，所以只需要用一个cap为2的切片即可。（-1表示篮子为空）

这题的关键在于，**怎么处理好滑动窗口的收缩（即left的更新问题）**。当不满足滑动窗口条件时，我们先将滑动窗口极致的收缩（`left=i`），然后向左扩大滑动窗口（`left--`）直到不满足条件。

其实搞清楚了滑动窗口思想，这题就很简单了。参考如下草稿：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221008114752.png)



代码如下：

```go
func totalFruit(fruits []int) int {
	res := 0
	fruitType := []int{-1, -1} // 表示两个篮子存放水果的种类(-1表示篮子中没有水果)
	left, i := 0, 0
	for i < len(fruits) {
		// 如果遍历到的水果种类在篮子中，则滑动窗口向右扩展（i++）
		if fruits[i] == fruitType[0] || fruits[i] == fruitType[1] {
			res = max(res, i-left+1)
			i++
			continue
		}
		// 考虑到篮子可能没有放水果的情况
		if fruitType[0] == -1 || fruitType[1] == -1 {
			if fruitType[0] == -1 {
				fruitType[0] = fruits[i] // 将遍历到的水果放入到篮子中
			} else {
				fruitType[1] = fruits[i] // 将遍历到的水果放入到篮子中
			}
			res = max(res, i-left+1)
			i++
			continue
		}
		// 更新left（收缩滑动窗口）
		left = i
		fruitType = []int{fruits[i], -1} // 滑动窗口更新了，更新篮子中的水果种类
		for {                            // 开始向左扩大滑动窗口
			// 如果遍历到的水果种类在篮子中，则滑动窗口向左扩展（left--）
			if fruits[left] == fruitType[0] || fruits[left] == fruitType[1] {
				res = max(res, i-left+1)
				left--
				continue
			}
			// 考虑到篮子可能没有放水果的情况
			if fruitType[1] == -1 { // 这里只需要考虑fruitType[1]即可（因为我们在上面更新篮子时是设置第二个篮子为空的）
				fruitType[1] = fruits[left] // 将遍历到的水果放入到篮子中
				res = max(res, i-left+1)
				left--
				continue
			}
			left++ // 滑动窗口不满足条件了，此时left++回到满足条件的情况
			break
		}
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
```

