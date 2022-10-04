# 283. 移动零

[LeetCode 283. 移动零](https://leetcode.cn/problems/move-zeroes/)

## 双指针

### 两次遍历

思路还是和之前一样，通过快慢指针`slow`、`fast`删除0元素：

- 第一次遍历：`fast`走到数组末尾，删除了整个数组的0元素
- 第二次遍历：第一次遍历结束后，slow在“非0新数组”的下一位，所以此时遍历slow指针，将后面的元素置0即可。

思路很简单打一下草稿就能搞明白。

代码如下：

```go
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
```

### 一次遍历

参考[王尼玛](https://leetcode.cn/u/wang_ni_ma/)的[题解](https://leetcode.cn/problems/move-zeroes/solution/dong-hua-yan-shi-283yi-dong-ling-by-wang_ni_ma/)。

这里参考了快速排序的思想，快速排序首先要确定一个待分割的元素做中间点x，然后把所有小于等于x的元素放到x的左边，大于x的元素放到其右边。

这里我们可以用0当做这个中间点，**把不等于0的放到中间点的左边，等于0的放到其右边**。

这里的中间点就是0本身，所以实现起来比快速排序简单很多，我们使用两个指针`slow`和`fast`：

- 开始遍历`fast`，`slow`不动
- 如果`nums[fast]!=0`，就交换`nums[fast]`和`nums[slow]`，然后slow向后移动一位（`slow++`）

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221004105033.gif)

代码如下：

```go
func moveZeroes(nums []int) {
	slow, fast := 0, len(nums)-1
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}
```



















