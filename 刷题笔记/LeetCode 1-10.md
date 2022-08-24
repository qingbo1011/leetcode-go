# 1.两数之和

[LeetCode 1. 两数之和](https://leetcode.cn/problems/two-sum/)

## 哈希表

具体实现不用多说了，之前写java的时候就实现过了。

代码如下：

```go
func twoSum(nums []int, target int) []int {
   res := make([]int, 0)
   m := make(map[int]int, 0) // key为nums[i]，value为i（即索引）
   for i, num := range nums {
      if value, ok := m[target-num]; ok {
         res = append(res, value, i)
         break
      }
      m[num] = i
   }
   return res
}
```



# 2. 两数相加

[LeetCode 2. 两数相加](https://leetcode.cn/problems/add-two-numbers/)

## 模拟

这是我按照自己的思路去写的：先依次遍历单链表，然后将每一位的结果存储到切片中（暂时不用考虑进位），接着统一处理进位得到最终的切片，最后将切片的值赋值给链表。

逻辑很简单，代码如下：

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	num := make([]int, 0)
	for l1 != nil && l2 != nil {
		num = append(num, l1.Val+l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}
	// 别忘记考虑l1和l2长度不一致的情况
	for l1 != nil {
		num = append(num, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		num = append(num, l2.Val)
		l2 = l2.Next
	}
	// 进位处理
	carry := 0
	for i, _ := range num {
		num[i] = num[i] + carry
		num[i], carry = num[i]%10, num[i]/10
	}
	if carry > 0 {
		num = append(num, carry)
	}
	// 进位处理之后就可以赋值给链表了
	temp := res
	for _, v := range num {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}
	return res.Next
}
```

> 在自己写的代码通过之后，参考[官方题解](https://leetcode.cn/problems/add-two-numbers/solution/liang-shu-xiang-jia-by-leetcode-solution/)，又整理了一下。
>
> 代码如下：
>
> ```go
> func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
>    head := &ListNode{} // 注意头结点和首元结点的区别
>    temp := head
>    carry := 0 // 进位
>    // 使用||，一次循环搞定l1，l2不对齐的情况
>    for l1 != nil || l2 != nil {
>       n1, n2 := 0, 0
>       // 在for循环内注意判断l1，l2是否已经到尽头了
>       if l1 != nil {
>          n1 = l1.Val
>          l1 = l1.Next
>       }
>       if l2 != nil {
>          n2 = l2.Val
>          l2 = l2.Next
>       }
>       sum := n1 + n2 + carry
>       sum, carry = sum%10, sum/10 // 得出当前位的结果，以及进位
>       temp.Next = &ListNode{Val: sum}
>       temp = temp.Next
>    }
>    if carry > 0 {
>       temp.Next = &ListNode{Val: carry}
>    }
>    return head.Next
> }
> ```

# 3. 无重复字符的最长子串

[LeetCode 3.无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

## 哈希表

最最进本的解法，哈希表。因为题目要求的是无重复最长子串的长度，而不是子串内容。就简单很多了。

使用map，**其中key存储字符，value存储字符在字符串中的索引**。

当发现重复字符时（假设两个重复字符a1，a2），就要从第一个重复字符（a1）的下一个字符继续遍历。

逻辑比较简单，打下草稿就能懂。

代码如下：

```go
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	m := make(map[byte]int, len(s)) // key为字符，value为该字符在字符串s中的索引
	tmp := 0
	for i := 0; i < len(s); i++ {
		if value, ok := m[s[i]]; ok { // map中已存在
			m = make(map[byte]int, len(s)) // 直接清空map
			ans = max(ans, tmp)
			tmp = 0
			i = value
		} else { // map中不存在
			m[s[i]] = i
			tmp++
		}
	}
	return max(ans, tmp)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

> 结果为：`解答成功:执行耗时:1416 ms,击败了5.33% 的Go用户。内存消耗:8.2 MB,击败了5.15% 的Go用户`

这里要注意的是关于go的map的清空：**Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map**，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。

## 滑动窗口

其实说是滑动窗口，实际上更像是双指针，滑动窗口只是一种形象的比喻。

因为这题要求的是最长子串的长度而不是内容，所以我们可以用类似双指针的思想来求出长度，其中`left`指向子串左边，遍历过程中的`i`指向子串右边，子串长度即为`i-left+1`。这样就不用再创建一个双端队列来存储子串了。

> - 首先，判断当前字符是否包含在map中，如果不包含，将该字符添加到map（key为字符，value为该字符在字符串中最后出现的索引（遍历过程中）），此时没有出现重复的字符，左指针不需要变化。此时不重复子串的长度为：`i-left+1`，与原来的max比较，取最大值。
> - 如果当前字符 ch 包含在 map中，此时有2类情况：
>   - **当前字符包含在当前有效的子段中**，如：`abca`，当我们遍历到第二个a，当前有效最长子段是 abc，我们又遍历到a，那么此时更新 `left` 为 `map[a]+1`（即我们在上面的方法中分析的，**一旦发现重复值，那么就从前一个重复元素之后的元素开始继续遍历**），当前有效子段更新为 bca；
>   - **当前字符不包含在当前最长有效子段中**，如：`abba`，我们先添加a,b进map，此时left=0，我们再添加b，发现map中包含b，而且b包含在最长有效子段中，我们更新 `left=map[b]+1`（=2），此时子段更新为 b，而且**map中仍然包含a**，map[a]=0；随后，我们遍历到a，发现a包含在map中，且map[a]=0，如果我们像上面那样处理，就会发现 left=map[a]+1=1，**实际上，left此时应该不变，left始终为2，子段变成 ba才对。**
>   - 为了解决上面两种情况的差异，我们在更新left时，**取left当前值和重复元素的索引+1这两者之中的最大值**！且**不管是否更新left，都要更新当前遍历字符的位置！**（因为我们的value是该字符在字符串中最后出现的索引），注意顺序是先去left再更新索引。

感觉这种处理，其实就是，免去了我们之前清空map的操作。

逻辑打下草稿就很清楚了。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220824212223.jpg)

代码如下：

```go
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int, 0)
	res := 0
	left := 0
	for i := range s {
		if index, ok := m[s[i]]; ok { // 字符重复
			left = max(left, index+1) // 新的left值取left当前值和重复元素的索引+1这两者之中的最大值
		}
		m[s[i]] = i // 不管是否更新left，都要更新当前遍历字符的位置！
		res = max(res, i-left+1)
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
```

> - 执行用时：4 ms, 在所有 Go 提交中击败了90.40%的用户
>
> - 内存消耗：2.9 MB, 在所有 Go 提交中击败了43.08%的用户























