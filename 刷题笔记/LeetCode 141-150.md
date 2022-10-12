# 141. 环形链表

[LeetCode 141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/?favorite=2cktkvj)

## 快慢指针

在一般面试时，比较经典的回答还是**双指针**。思路就是：使用两个指针`p1`，`p2`。**p1指针每次走一步，p2指针每次走两步**。在指针走到尽头之前，如果p1等于p2了，那么就说明有环。（事实上如果有环了话，两个指针就永远不刽走到尽头，迟早会相遇的）

代码如下：

```go
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head // 慢指针
	fast := head // 快指针
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```



# 142. 环形链表 II

[LeetCode 142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/?favorite=2cktkvj)

## 哈希表

最简单易懂的方法了。

代码如下：

```go
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	p := head
	for p != nil {
		if _, ok := m[p]; ok { // 找到重复索引了
			return p
		}
		m[p] = 1
		p = p.Next
	}
	return nil
}
```

> - 执行用时：8 ms, 在所有 Go 提交中击败了29.52%的用户
> - 内存消耗：5.8 MB, 在所有 Go 提交中击败了5.02%的用户

## 快慢指针

[官方题解](https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/)

**快慢指针，当slow指针和fast指针相遇后，再让ptr指针从头结点开始，和slow指针每次向后移动一个位置。最终，它们会在入环点相遇。**

> 数学推论：
>
> 我们使用两个指针，fast 与 slow。它们起始都位于链表的头部。随后，slow指针每次向后移动一个位置，而 fast 指针向后移动两个位置。如果链表中存在环，则 fast 指针最终将再次与 slow 指针在环中相遇。
>  如下图所示，设链表中环外部分的长度为 a。 slow 指针进入环后，又走了 b 的距离与 fast 相遇。此时，fast 指针已经走完了环的 n 圈，因此它走过的总距离为 a+n(b+c)+b=a+(n+1)b+nc
>
> [![img](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220502151449.png)](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220502151449.png)
>
> 根据题意，任意时刻，fast 指针走过的距离都为 slow 指针的 2 倍。因此，我们有
>
> a+(n+1)b+nc=2(a+b)⟹a=c+(n−1)(b+c)
>
> 有了 a=c+(n−1)(b+c) 的等量关系，我们会发现：从相遇点到入环点的距离加上 n-1 圈的环长，恰好等于从链表头部到入环点的距离。
>
> **因此，当发现 slow 与 fast 相遇时，我们再额外使用一个指针 ptr。起始，它指向链表头部；随后，它和slow每次向后移动一个位置。最终，它们会在==入环点==相遇。**

代码如下：

```go
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	pre := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 先移动在判断，因为slow和fast初始都在head上
		if slow == fast {
			for slow != pre {
				slow = slow.Next
				pre = pre.Next
			}
			return pre
		}
	}
	return nil
}
```

> - 执行用时：4 ms, 在所有 Go 提交中击败了94.83%的用户
> - 内存消耗：3.5 MB, 在所有 Go 提交中击败了27.27%的用户























