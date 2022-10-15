# 202. 快乐数

[LeetCode 202. 快乐数](https://leetcode.cn/problems/happy-number/)

## HashSet

> [参考官方题解](https://leetcode.cn/problems/happy-number/solution/kuai-le-shu-by-leetcode-solution/)
>
> 根据我们的探索，我们猜测会有以下三种可能：
>
> 1. 最终会得到1
> 2. 最终会进入循环
> 3. 值会越来越大，最后接近无穷大（后面会说明这种情况永远不会发生）
>
> 第三个情况比较难以检测和处理。我们怎么知道它会继续变大，而不是最终得到 11 呢？我们可以仔细想一想，每一位数的最大数字的下一位数是多少。
>
> | Digits |    Largest    | Next |
> | ------ | :-----------: | ---: |
> | 1      |       9       |   81 |
> | 2      |      99       |  162 |
> | 3      |      999      |  243 |
> | 4      |     9999      |  324 |
> | 13     | 9999999999999 | 1053 |
>
> 对于3位数的数字，它不可能大于243。这意味着它要么被困在243以下的循环内，要么跌到1。4位或4 位以上的数字在每一步都会丢失一位，直到降到3位为止。所以我们知道，最坏的情况下，算法可能会在 243 以下的所有数字上循环，然后回到它已经到过的一个循环或者回到1。但它不会无限期地进行下去，所以我们**排除第三种选择**。
>
> 即使在代码中你不需要处理第三种情况，你仍然需要理解为什么它永远不会发生，这样你就可以证明为什么你不处理它。
>

这题乍一看比较麻烦，但是只要获取了题目所给的关键信息：然后重复这个过程**直到这个数变为1**，也可能是**无限循环**，但始终变不到1。

也就是说，给定的数在循环（重复的这个过程）中只会有两种可能：

- 最终结果为1，`return true`
- 陷入无线循环，即过程中出现了重复的数字，`return false`

那么代码只要就只需要解决两个问题了：

- 判断是否出现了重复的数字，直接使用set。

- 获取一个数每一次将该数替换为它每个位置上的数字的平方和的数，我们定义一个`getSun`函数，代码如下：

  ```go
  func getSum(n int) int {
     sum := 0
     for n > 0 {
        sum = sum + (n%10)*(n%10) // n最后一位进行平方
        n = n / 10                // n去除最后一位
     }
     return sum
  }
  ```

完整代码如下：

```go
func isHappy(n int) bool {
	set := make(map[int]struct{}, 0)
	for {
		if n == 1 {
			return true
		}
		n = getSum(n)
		if _, ok := set[n]; ok { // 出现了重复，说明陷入无限循环了
			return false
		}
		set[n] = struct{}{}
	}
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum = sum + (n%10)*(n%10) // n最后一位进行平方
		n = n / 10                // n去除最后一位
	}
	return sum
}
```



# 203. 移除链表元素

[LeetCode 203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/)

## 单链表

经典的删除单链表的结点。不用多说什么，主需要注意两点：

- dummyNode的使用

- **删除一个结点，那么我们要遍历到当前结点的前一个结点**

  ![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221011202655.jpg)

代码如下：

```go
func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Next: head}
	tmp := dummyNode
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val { // tmp.Next是要删除的节点
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyNode.Next
}
```

（结合代码画一下草稿就搞懂了）



# 206. 反转链表

[LeetCode 206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)

## 单链表

之前用Java写过，具体逻辑就不细说了。根据代码打下草稿就能懂。

代码如下：

```go
func reverseList(head *ListNode) *ListNode {
   var pre *ListNode
   cur := head
   for cur != nil {
      next := cur.Next
      cur.Next = pre // 修改结点指向
      pre = cur
      cur = next
   }
   return pre
}
```



# 209. 长度最小的子数组

[LeetCode 209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)

## 滑动窗口

使用滑动窗口的思想（还是类似双指针）。

使用文字说明可能会看起来更复杂，直接看分析草稿：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220825160551.jpg)

> 这个题解的滑动窗口解法和我想的一样：[官方题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solution/chang-du-zui-xiao-de-zi-shu-zu-by-leetcode-solutio/)

代码如下：

```go
func minSubArrayLen(target int, nums []int) int {
	//length := len(nums)
	res := math.MaxInt
	left := 0
	sum := 0
	for i, num := range nums {
		sum = sum + num
		for sum >= target {
			res = min(res, i-left+1)
			sum = sum - nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```





















