# 234. 回文链表

## 双指针

这题我一开始想的是，将链表里的val存储到`[]int`切片里，然后通过双指针判断是否是回文。

思路很简单，代码如下：

```go
func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```



















