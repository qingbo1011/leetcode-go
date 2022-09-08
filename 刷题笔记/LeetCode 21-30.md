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



# 30. 串联所有单词的子串

[LeetCode 30. 串联所有单词的子串](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/)

## 滑动窗口

这题是[LeetCode 438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)的进阶版，可以先参考[LeetCode 438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)的题解再来看此题。



