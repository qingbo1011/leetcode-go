# 704. 二分查找

[LeetCode 704. 二分查找](https://leetcode.cn/problems/binary-search/)

## 二分查找

这是一道经典的二分查找题目。

二分查找的思路很简单，但是代码写起来却很容易犯错。主要体现在两个方面：

- 例如到底是 `for left < right` 还是 `for left <= right`
- 到底是`right = middle`呢，还是要`right = middle - 1`呢

区分上面两种方法，主要是看区间是怎么定义的。写二分法，区间的定义一般为两种，**左闭右闭**即`[left, right]`，或者左闭右开即[left, right)。

这里我们给出**左闭右闭**的区间，即：

- `for left <= right`
- `left = middle + 1`
- `right = middle - 1`

> 如果是左闭右开即`[left, right)`：
>
> - `for left < right`
> - `left = middle + 1`
> - `right = middle`

代码如下：

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
```



# 707. 设计链表

[LeetCode 707. 设计链表](https://leetcode.cn/problems/design-linked-list/)

## 单链表

[官方题解](https://leetcode.cn/problems/design-linked-list/solution/she-ji-lian-biao-by-leetcode-solution-abix/)

数据结构采用单向链表。

- `MyLinkedList`中的成员有：

  - **哨兵结点dummy**（这个不用过多介绍了，单链表题里很常见了）
  - 当前链表的节点数size

- 当然我们要新建一个数据结构`Node`，成员有：

  - 当前节点值Val
  - 当前节点指向的下一个节点Next

  那么定义数据结构的代码就如下：

  ```go
  type MyLinkedList struct {
  	dummy *Node
  	size  int
  }
  
  type Node struct {
  	Val  int
  	Next *Node
  }
  ```

- 定义好了数据结构，构造函数`Constructor`就比较好写了：

  ```go
  func Constructor() MyLinkedList {
  	return MyLinkedList{dummy: &Node{}, size: 0}
  }
  ```

- get(index)：获取链表中第 `index` 个节点的值。如果索引无效，则返回`-1`。

  因为`MyLinkedList`中有`size`这个成员变量，题目中也说了链表中的所有节点都是 0-index 的，所以就比较好写了。代码如下：

  ```go
  func (l *MyLinkedList) Get(index int) int {
     if index < 0 || index >= l.size { // index不合法
        return -1
     }
     cur := l.dummy
     for i := 0; i <= index; i++ { // 因为cur是从哨兵节点开始的
        cur = cur.Next
     }
     return cur.Val
  }
  ```

- `addAtHead(val)`：在链表的第一个元素之前添加一个值为 `val` 的节点。插入后，新节点将成为链表的第一个节点。

  这个没什么好说的，写好`addAtIndex(index,val)`即可，然后在内部执行`l.AddAtIndex(0, val)`。代码如下：

  ```go
  func (l *MyLinkedList) AddAtHead(val int) {
     l.AddAtIndex(0, val)
  }
  ```

  （这里不`l.size++`是因为在`addAtIndex(index,val)`中已经`l.size++`了）

- `addAtTail(val)`：将值为 `val` 的节点追加到链表的最后一个元素。

  这个也没什么好说的，写好`addAtIndex(index,val)`即可，然后在内部执行`l.AddAtIndex(l.size, val)`。代码如下：

  ```go
  func (l *MyLinkedList) AddAtTail(val int) {
     l.AddAtIndex(l.size, val)
  }
  ```

  （这里不`l.size++`是因为在`addAtIndex(index,val)`中已经`l.size++`了）

- addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。

  处理一下边界条件，就可以开始链表的添加节点操作了。要在某一个节点前添加节点，我们要先找到该节点的前一个节点，即**要在cur节点前添加一个节点，我们要找到cur节点的前一个节点pre**（这个画一下图就可以理解了）。最后别忘了`l.size++`。

  代码如下：

  ```go
  func (l *MyLinkedList) AddAtIndex(index int, val int) {
     if index > l.size { // 如果index大于链表长度，则不会插入节点
        return
     }
     if index < 0 { // 如果index小于0，则在头部插入节点
        index = 0
     }
     // 接下来就是链表的添加操作了
     pre := l.dummy
     for i := 0; i < index; i++ {
        pre = pre.Next
     }
     cur := pre.Next
     pre.Next = &Node{Next: cur, Val: val}
     l.size++ // 别忘了链表size+1
  }
  ```

- deleteAtIndex(index)：如果索引 `index` 有效，则删除链表中的第 `index` 个节点。

  首先还是处理好边界条件（index不合法的情况），然后就是经典的单链表删除节点了。跟添加节点一样，**要删除cur节点，我们要找到cur节点的前一个节点pre**。最后别忘了`l.size--`。

完整代码如下：

```go
type MyLinkedList struct {
	dummy *Node
	size  int
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &Node{}, size: 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size { // index不合法
		return -1
	}
	cur := l.dummy
	for i := 0; i <= index; i++ { // 因为cur是从哨兵节点开始的
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size { // 如果index大于链表长度，则不会插入节点
		return
	}
	if index < 0 { // 如果index小于0，则在头部插入节点
		index = 0
	}
	// 接下来就是链表的添加操作了
	pre := l.dummy
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	pre.Next = &Node{Next: cur, Val: val}
	l.size++ // 别忘了链表size+1
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size { // index不合法
		return
	}
	// 开始删除节点
	pre := l.dummy
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	l.size-- // 别忘了链表size-1
}
```

 













































