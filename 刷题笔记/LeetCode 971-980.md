# 977. 有序数组的平方

[LeetCode 977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/)

## 直接排序

方法很简单，将原数组平方处理后直接对新数组排序。

代码如下：

```go
func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	return quicksort(nums)
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := 0, len(nums)-1
	pivot := 0
	nums[pivot], nums[right] = nums[right], nums[pivot]
	for i := range nums {
		if nums[i]<nums[right] {
			nums[left],nums[i] = nums[i],nums[left]
			left++
		}
	}
	nums[left],nums[right] = nums[right],nums[left]

	quicksort(nums[:left])
	quicksort(nums[left+1:])

	return nums
}
```

> 或者直接用Go的sort包。
>
> ```go
> func sortedSquares(nums []int) []int {
> 	for i := range nums {
> 		nums[i] = nums[i] * nums[i]
> 	}
> 	sort.Ints(nums)
> 	return nums
> }
> ```

## 双指针

**进阶：**请你设计时间复杂度为 `O(n)` 的算法解决本问题。

[官方题解](https://leetcode.cn/problems/squares-of-a-sorted-array/solution/you-xu-shu-zu-de-ping-fang-by-leetcode-solution/)（方法三）

上面的直接排序，并没有利用**数组nums按升序排序**这个条件。

**先新建一个答案数组，在遍历过程中依次将结果逆序的放入答案数组中**。我们可以使用两个指针分别指向位置 0 和 n-1，每次比较两个指针对应的数，选择较大的那个**逆序**放入答案并移动指针。这种方法无需处理某一指针移动至边界的情况，因为我们是在`for i := len(nums) - 1; i >= 0; i--`的遍历过程中依次赋值。可以结合代码自行体会。

代码如下：

```go
func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums)) // 创建一个len为len(nums)的答案切片
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if v, k := nums[left]*nums[left], nums[right]*nums[right]; v < k { // 右边的平方大
			ans[i] = k
			right--
		} else {
			ans[i] = v
			left++
		}
	}
	return ans
}
```

> - 时间复杂度：O(n)，其中 n 是数组nums 的长度。
>
> - 空间复杂度：O(1)。除了存储答案的数组以外，我们只需要维护常量空间。
>



















