# 19. 删除链表的倒数第 N 个结点

[LeetCode 19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

## 双指针

按照之前用Java写过的逻辑，求链表的的倒数第N个结点，使用双指针即可：

- （因为涉及到删除，所以要让p1走到倒数第N个结点的前一个，这样更好删除）

- 初始p1，p2都为`dummyNode`（dummyNode.Next = head），p2指针永远比p1指针快N个结点，然后同时移动p1和p2，这样当`p2.Next==nil`时，p1指向的即为链表的倒数第N个结点的前一个结点
- 找到了倒数第N个结点的前一个p1，要删除倒数第N个就很简单了（`p1.Next = p1.Next.Next`）

思路很简单画一下图就能懂，代码如下：

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	for i := 0; i < n; i++ { // 双指针就位
		p2 = p2.Next
	}
	for p2.Next != nil { // p1就位到要删除节点的前一个位置
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next // 删除节点
	return dummy.Next
}
```



# 20. 有效的括号

[LeetCode 20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/)

使用map作为左右括号的匹配，其中：

- **key为右括号**
- **value为左括号**

> 这里顺序很重要不能搞反了。

然后我们创建一个切片用来模拟栈。当栈中的**栈底元素是一个右括号**时，就不用继续遍历了，此时不可能是有效的括号（这个很好理解）。

接下来的逻辑就是遍历匹配了，画一下草图就能理解，代码如下：

```go
func isValid(s string) bool {
   length := len(s)
   if length%2 != 0 {
      return false
   } // 奇数长度就不可能为有效括号了
   stack := make([]byte, 0) // 自定义的栈
   for i := 0; i < length; i++ {
      if len(stack) > 0 { // 栈中有元素
         if stack[0] == ')' || stack[0] == ']' || stack[0] == '}' { // 栈底是右括号，直接返回false
            return false
         }
         // 匹配到一对括号，栈顶元素出栈
         if (stack[len(stack)-1] == '(' && s[i] == ')') ||
            (stack[len(stack)-1] == '[' && s[i] == ']') ||
            (stack[len(stack)-1] == '{' && s[i] == '}') {
            stack = stack[:len(stack)-1]
         } else { // 未匹配到一对括号，新元素入栈
            stack = append(stack, s[i])
         }
      } else { // 栈为空，元素直接进栈
         stack = append(stack, s[i])
      }
   }
   return len(stack) == 0
}
```



