# 21. 合并两个有序链表

[LeetCode 21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/?favorite=2cktkvj)

## 双指针

合并两个有序列表，就跟合并两个有序数组差不多。使用双指针即可。直接看代码吧。

代码如下：

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	tmp := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = list2
			tmp = tmp.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		tmp.Next = list1
		tmp = tmp.Next
		list1 = list1.Next
	}
	for list2 != nil {
		tmp.Next = list2
		tmp = tmp.Next
		list2 = list2.Next
	}
	return res.Next
}
```



# 27. 移除元素

[LeetCode 27. 移除元素](https://leetcode.cn/problems/remove-element/)

## 双指针

[参考代码随想录： 移除元素](https://programmercarl.com/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.html#%E6%80%9D%E8%B7%AF)

定义快慢指针：

- 快指针：寻找新数组的元素 （新数组就是不含有目标元素的数组）
- 慢指针：指向更新新数组下标的位置

> 快指针fast向右遍历旧数组：
>
> - 如果没有遇到要删除的元素，该元素是新数组的成员。所以将新元素赋值给慢指针slow的位置，并向右移动慢指针slow（）
> - 如果遇到了要删除的元素，该元素不是新数组的成员。**此时什么都不用做，直接移动快指针fast即可**。（因为不需要更新慢指针slow）

删除过程如下：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220916205643.gif)

代码如下：

```go
func removeElement(nums []int, val int) int {
	length := len(nums)
	slow, fast := 0, 0
	for fast = 0; fast < length; fast++ {
		if nums[fast] != val { // 是新数组成员，要更新慢指针
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow] // 更新一下数组
	return slow
}
```

> - 时间复杂度：O(n)
> - 空间复杂度：O(1)



# 30. 串联所有单词的子串

[LeetCode 30. 串联所有单词的子串](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/)

## 滑动窗口

这题是[LeetCode 438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)的进阶版，可以先参考[LeetCode 438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)的题解再来看此题。



