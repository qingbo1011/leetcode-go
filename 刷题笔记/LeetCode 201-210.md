# 206. 反转链表

[LeetCode 206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)

## 单链表

之前用Java写过，具体逻辑就不细说了。根据代码打下草稿就能懂。

代码如下：

```go
func reverseList(head *ListNode) *ListNode {
   var pre *ListNode
   cur := head
   for cur != nil {
      next := cur.Next
      cur.Next = pre // 修改结点指向
      pre = cur
      cur = next
   }
   return pre
}
```



# 209. 长度最小的子数组

[LeetCode 209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)

## 滑动窗口

使用滑动窗口的思想（还是类似双指针）。

使用文字说明可能会看起来更复杂，直接看分析草稿：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220825160551.jpg)

> 这个题解的滑动窗口解法和我想的一样：[官方题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solution/chang-du-zui-xiao-de-zi-shu-zu-by-leetcode-solutio/)

代码如下：

```go
func minSubArrayLen(target int, nums []int) int {
	//length := len(nums)
	res := math.MaxInt
	left := 0
	sum := 0
	for i, num := range nums {
		sum = sum + num
		for sum >= target {
			res = min(res, i-left+1)
			sum = sum - nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```





















