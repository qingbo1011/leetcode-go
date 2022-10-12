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



# 24. 两两交换链表中的节点

[LeetCode 24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/)

## 双指针

因为这题要求：你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。所以就不能取巧去交换Val值。

> 这题我其实是有点取巧的：因为是单向链表且无循环，所以交换两个节点的处理我在这里是**直接交换两个节点的Val值**的：`cur.Val, next.Val = next.Val, cur.Val`。这样处理一下其他的就很好写了，做一个循环，然后`cur`、`next`指针每次向后移动两位。
>
> 代码如下：
>
> ```go
> func swapPairs(head *ListNode) *ListNode {
> 	if head == nil {
> 		return nil
> 	}
> 	cur, next := head, head.Next
> 	for next != nil {
> 		cur.Val, next.Val = next.Val, cur.Val // 交换两节点的val
> 		cur = cur.Next.Next
> 		if next.Next == nil { // 做一下判断，不然next = next.Next.Next会panic
> 			break
> 		}
> 		next = next.Next.Next
> 	}
> 	return head
> }
> ```

不取巧，真正的交换节点的代码，参考[官方题解](https://leetcode.cn/problems/swap-nodes-in-pairs/solution/liang-liang-jiao-huan-lian-biao-zhong-de-jie-di-91/)。

- `temp=dummy`（哨兵节点dummy，这种处理不用再多说）
- 每次循环中，`pre=temp.Next`，`next=temp.Next.Next`
- 要交换pre和next，那么就需要temp参与其中，3个指针来处理：`temp.Next=next`、`pre.Next=next.Next`、`next.Next=pre`（这个逻辑打一下草稿就能懂）
- 开始下一次循环，移动temp：`temp=pre`（打下草稿就很清楚了）
- 维持循环的条件：`temp.Next!=nil && temp.Next.Next!=nil`（这个取一下节点奇偶数的特殊情况就能搞明白了，不需要加`temp!=nil`）

```go
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		// 双指针就位
		pre := temp.Next
		next := temp.Next.Next
		// 利用temp、pre、next这三个指针开始交换pre和next这两个节点
		temp.Next = next
		pre.Next = next.Next
		next.Next = pre
		// 交换完成，更新temp继续循环
		temp = pre
	}
	return dummy.Next
}
```



# 26. 删除有序数组中的重复项

[LeetCode 26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

## 双指针

参考题解：[【双指针】删除重复项-带优化思路](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/solution/shuang-zhi-zhen-shan-chu-zhong-fu-xiang-dai-you-hu/)

注意**数组是有序的，那么重复的元素一定会相邻**。要求删除重复元素，实际上就是**将不重复的元素移到数组的左侧**。

所以我们还是定义快慢指针slow和fast（初始slow和fast都为0），比较slow和fast位置的元素是否相等：

- 如果相等（`nums[slow]==nums[fast]`）：slow不变，fast后移一位（`fast++`）
- 如果不相等（`nums[slow]!=nums[fast]`）：将fast位置的元素复制到slow的**后一位**（`nums[slow+1]=nums[fast]`），然后slow和fast都后移一位（`slow++`，`fast++`）

fast遍历到原数组的最后一位时循环结束。**循环结束后slow指针位于新数组的最后一位**，即**删除重复元素的新数组长度为slow+1**。

> ![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221003112039.png)

代码如下：

```go
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast = 0; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow+1] // 更新数组
	return slow + 1
}
```



# 27. 移除元素

[LeetCode 27. 移除元素](https://leetcode.cn/problems/remove-element/)

## 双指针

- [代码随想录： 移除元素](https://programmercarl.com/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.html#%E6%80%9D%E8%B7%AF)
- **[B站视频](https://www.bilibili.com/video/BV12A4y1Z7LP/)**

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

> 参考[powcai](https://leetcode.cn/u/powcai/)的[题解](https://leetcode.cn/problems/substring-with-concatenation-of-all-words/solution/chuan-lian-suo-you-dan-ci-de-zi-chuan-by-powcai/)下[时间是个贼](https://leetcode.cn/u/zhang-zhi-qiang-5/)提供的golang代码：
>
> ```go
> var oriMap map[string]int
> 
> func findSubstring(s string, words []string) []int {
> 	oriMap = map[string]int{}
> 	for _, w := range words {
> 		oriMap[w]++
> 	}
> 	wordSize := len(words[0])
> 	stepSize := wordSize * len(words)
> 	var check = func(sk string) bool {
> 		nowMap := map[string]int{}
> 		for i := 0; i+wordSize <= len(sk); i += wordSize {
> 			wk := sk[i : i+wordSize]
> 			nowMap[wk]++
> 		}
> 		return reflect.DeepEqual(nowMap, oriMap)
> 	}
> 	res := []int{}
> 	for i := 0; i+stepSize <= len(s); i++ {
> 		ss := s[i : i+stepSize]
> 		if check(ss) {
> 			res = append(res, i)
> 		}
> 	}
> 	return res
> }
> ```

这题运用滑动窗口的思想，维护一个符合条件的滑动窗口即可。需要注意的是，意图给定了`words` 中所有字符串 **长度相同**（`wordSize`）。所以我们**的滑动窗口的长度一定是固定的！**（`stepSize := wordSize * len(words)`）。

- 所以我们遍历i，然后通过`windowString := s[i : i+stepSize]`获取到滑动窗口内的字符串
- 接着，我们将滑动窗口类的字符串统计成`map`，通过`reflect.DeepEqual`的方法判断滑动窗口内的字符串是否符合条件。若符合则append到结果中；不符合就继续遍历i。

搞懂了思想这题就没这么难了（主要是因为通过`reflect.DeepEqual`来判断两个map的方式，判断滑动窗口是否符合条件，这样就大大降低了难度）

代码如下：

```go
func findSubstring(s string, words []string) []int {
	var res []int
	wordSize := len(words[0])           // 单个单词的长度
	stepSize := wordSize * len(words)   // 滑动窗口的长度
	wordsMap := make(map[string]int, 0) // words转成map，与滑动窗口中的windowMap进行比较
	for _, word := range words {        // 根据words获取到wordsMap
		wordsMap[word]++
	}
	// 开始遍历滑动窗口
	for i := 0; i+stepSize-1 < len(s); i++ { // i+stepSize-1即为滑动窗口最后一个元素所在索引
		windowString := s[i : i+stepSize] // [i,i+stepSize-1]，即为当前遍历到的滑动窗口内的字符串
		if check(windowString, wordsMap, wordSize) {
			res = append(res, i)
		}
	}
	return res
}

// 根据windowString和wordsMap，来判断当前的滑动窗口内的字符串是否符合条件（传入wordSize以便将windowString转为wordsMap）
func check(windowString string, wordsMap map[string]int, wordSize int) bool {
	// 首先将字符串windowString同样的转为map
	windowMap := make(map[string]int, 0)
	for i := 0; i+wordSize-1 < len(windowString); i = i + wordSize { // 看起来复杂其实很简单
		windowMap[windowString[i:i+wordSize]]++
	}
	return reflect.DeepEqual(windowMap, wordsMap)
}
```

> 参考：
>
> - [Go reflect.DeepEqual](https://www.jianshu.com/p/50380740d04a)
> - [reflect.DeepEqual函数：判断两个值是否一致](https://blog.csdn.net/m0_37710023/article/details/108284171)
> - [reflect.DeepEqual() Function in Golang with Examples](https://www.geeksforgeeks.org/reflect-deepequal-function-in-golang-with-examples/)
>
> 这里补充一下go中的`reflect.DeepEqual`：The **reflect.DeepEqual()** Function in Golang is used to check whether x and y are “deeply equal” or not.
>
> 先直接看源码及源码中的部分注释：
>
> ```go
> // DeepEqual reports whether x and y are ``deeply equal,'' defined as follows.
> // Two values of identical type are deeply equal if one of the following cases applies.
> // Values of distinct types are never deeply equal.
> // ...
> // Map values are deeply equal when all of the following are true:
> // they are both nil or both non-nil, they have the same length,
> // and either they are the same map object or their corresponding keys
> // (matched using Go equality) map to deeply equal values.
> func DeepEqual(x, y any) bool {
>    if x == nil || y == nil {
>       return x == y
>    }
>    v1 := ValueOf(x)
>    v2 := ValueOf(y)
>    if v1.Type() != v2.Type() {
>       return false
>    }
>    return deepValueEqual(v1, v2, make(map[visit]bool))
> }
> ```
>
> 对于`slice`、`map`、`struct`等类型，当比较两个值是否相等时，是不能使用`==`运算符的。

这里偷懒就直接使用`reflect.DeepEqual`了。如果要优化的话就需要自己手写一个判断滑动窗口是否合法的函数：

- 首先比较两个map中的key是否一样
- 其次比较map中的key锁对应的value是否一样

























