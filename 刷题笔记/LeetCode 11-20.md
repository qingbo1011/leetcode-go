# 15. 三数之和

[LeetCode 15. 三数之和](https://leetcode.cn/problems/3sum/)

## 排序+双指针

三数之和不同于两数之和（[LeetCode 1. 两数之和](https://leetcode.cn/problems/two-sum/)），如果使用哈希表的方法会非常麻烦。所以这里我们使用**排序+双指针**的方法。

参考：

- [官方题解](https://leetcode.cn/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/)
- **[代码随想录B站视频](https://www.bilibili.com/video/BV1GW4y127qo/?)**

这题最关键的部分就在去重，其中具体细节都在两层循环的代码中。可以重点参考一下上面的题解和视频。

我们要找到`nums[i]+nums[left]+nums[right]==0`这三个数，同时注意对结果集res进行去重：

- 第一个数`nums[i]`的去重：确保`nums[i]`不会重复使用（即出现`nums[i]==nums[i-1]`，说明不能遍历当前的`i`，不然会重复）

  ```go
  if i > 0 && nums[i] == nums[i-1] { // 保证nums[i]不会在结果集中重复
     continue
  }
  ```

- 对后两个数nums[left]和nums[right]的去重：收缩后面的区间时，如果出现`nums[left] == nums[left+1]`，应该`left++`；如果出现`nums[right] == nums[right-1]`，应该`right--`

  ```go
  for left < right && nums[left] == nums[left+1] { // 对nums[left]进行去重
     left++
  }
  for left < right && nums[right] == nums[right-1] { // 对nums[right]进行去重
     right--
  }
  ```

代码如下：

```go
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums) // 先排序
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // 后面再也不可能出现三数之和为0的三元组了
			return res
		}
		if i > 0 && nums[i] == nums[i-1] { // 保证nums[i]不会在结果集中重复
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 { // 需要更新left
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 { // 需要更新right
				right--
			} else { // 找到结果集了
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] { // 对nums[left]进行去重
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 对nums[right]进行去重
					right--
				}
				// 去重后收缩left和right，让for left < right循环可以找到出口
				left++
				right--
			}
		}
	}
	return res
}
```



# 18. 四数之和

[LeetCode 18. 四数之和](https://leetcode.cn/problems/4sum/)

## 排序+双指针

参考：

- **[代码随想录视频](https://www.bilibili.com/video/BV1DS4y147US/?)**
- [官方题解](https://leetcode.cn/problems/4sum/solution/si-shu-zhi-he-by-leetcode-solution/)

四数之和，和[LeetCode 15. 三数之和](https://leetcode.cn/problems/3sum/)是一个思路，都是使用排序+双指针，基本解法就是在[LeetCode 15. 三数之和](https://leetcode.cn/problems/3sum/)的基础上再套一层for循环。

对于[LeetCode 15. 三数之和](https://leetcode.cn/problems/3sum/)双指针法就是将原本暴力O(n^3^)的解法，降为O(n^2^)的解法，四数之和的双指针解法就是将原本暴力O(n^4^)的解法，降为O(n^3^)的解法。



具体的逻辑可以看看上面的视频。

代码如下：













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

## 栈

我们创建一个切片用来模拟栈。当栈中的**栈底元素是一个右括号**时，就不用继续遍历了，此时不可能是有效的括号（这个很好理解）。

接下来的逻辑就是遍历匹配了，画一下草图就能理解。

代码如下：

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



