# 160. 相交链表

[LeetCode 160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/?favorite=2cktkvj)

## 哈希表

直接使用哈希表来判断是否有相交链表，思路简单暴力。

代码如下：

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
```

> 在Go中，空结构体`struct{}`是不占任何内存的。因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用。一是节省资源。二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符。（其中经典的用法就是用来实现集合Set。这里可以引申讲讲[Go空结构体struct{}](https://www.xiuxiubiu.com/archives/go%E7%A9%BA%E7%BB%93%E6%9E%84%E4%BD%93struct)）

## 双指针

使用双指针的方法，可以将**空间复杂度**降至 O(1)。

[官方题解](https://leetcode.cn/problems/intersection-of-two-linked-lists/solution/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/)

- 只有当链表headA 和 headB 都不为空时，两个链表才可能相交。因此首先判断链表 headA 和 headB 是否为空，如果其中至少有一个链表为空，则两个链表一定不相交，返回 null。
- 当链表headA 和 headB 都不为空时，创建两个指针pA 和pB，初始时分别指向两个链表的头节点 headA 和 headB，然后将两个指针依次遍历两个链表的每个节点。具体做法如下：
  - 每步操作需要同时更新指针 pA 和 pB。
  - 如果指针pA 不为空，则将指针 pA 移到下一个节点；如果指针 pB 不为空，则将指针 pB 移到下一个节点。
  - **如果指针 pA 为空，则将指针pA 移到链表headB 的头节点**；**如果指针pB 为空，则将指针 pB 移到链表 headA 的头节点**。
  - **当指针pA 和pB 指向同一个节点时，返回它们指向的节点**；**若指针pA 和pB 都为空，返回null**。

> ![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220907172458.png)

代码如下：

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
```





























