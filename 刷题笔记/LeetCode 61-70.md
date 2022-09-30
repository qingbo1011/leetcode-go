# 69. x 的平方根

[LeetCode 69. x 的平方根](https://leetcode.cn/problems/sqrtx/)

## 二分查找

这题的二分查找的边界条件是很容易找错的。

参考[liweiwei1419](https://leetcode.cn/u/liweiwei1419/)大佬的[题解](https://leetcode.cn/problems/sqrtx/solution/er-fen-cha-zhao-niu-dun-fa-python-dai-ma-by-liweiw/)。

首先根据题意：

> 示例 1：
>
> ```
> 输入: 4
> 输出: 2
> ```
>
> 这是显然的，4本身是一个完全平方数，2^2 = 4。虽然在数学上一个数的平方根有正有负，但是这个题目只要求我们返回算术平方根。
>
> 示例 2 ：
>
> ```
> 输入: 8
> 输出: 2
> ```
>
> 因为 8 的平方根实际上是 2.82842，题目要求我们将小数部分舍去。因此输出 2。于是我们知道：由于输出结果的时候，需要将小数部分舍去。

因此问题的答案，**平方以后一定不会严格大于输入的整数**。

思路分析：那么这就是一个查找整数的问题，并且这个整数是有范围的：

- 如果这个整数的平方 `==` 输入整数，那么我们就找到了这个整数；
- 如果这个整数的平方 `>` 输入整数，那么这个整数**肯定不是**我们要找的那个数；
- 如果这个整数的平方 `<` 输入整数，那么这个整数**==可能==**是我们要找的那个数（重点理解这句话，比如`输入: 8；输出: 2`）。

因此我们可以使用**二分查找**来查找这个整数，不断缩小范围去猜。

- **猜的数平方以后大了就往小了猜**（`right = middle - 1`）
- 猜的数平方以后恰恰好等于输入的数就找到了（`return middle`）
- **猜的数平方以后小了，可能猜的数就是，也可能不是**（`left = middle`）

代码如下：

```go
func mySqrt(x int) int {
	// 特殊值判断
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x/2
	for left < right {
		middle := (left + right + 1) >> 1
		if middle*middle == x {
			return middle
		} else if middle*middle < x { // 更新left（结果可能是middle）
			left = middle
		} else { // 更新right（结果不可能是right，因为都middle*middle大于x了）
			right = middle - 1
		}
	}
	return left
}
```

> 写 `if` 和 `else` 的原因：
>
> - `middle == x/middle`：`middle`即为结果，直接`return middle`
> - `middle > x/middle`：**`middle ` 肯定不是解，比 `middle ` 大的整数也肯定不是解**，因此问题的答案只可能存在区间 `[left, middle  - 1]`，所以设置 `right = middle - 1`
> - `middle < x/middle`：`middle > x/middle`情况的反例，既然大于的情况`[left, middle  - 1]`，那么`middle < x/middle`区间就应该是`[middle, right]`，所以设置`left = middle`
>
> 为什么最后返回 `left`：退出 `for left < right`循环的时候，由于边界搜索是 `left = middle` 与  `right = middle - 1`，**因此退出循环的时候一定有 left 与 right 重合**，**返回 `right` 也可以。**
>

**==问题==**：`middle` 为什么要加 `1`？

对着错误的测试用例打印出变量 `left` 、`right` 和 `mid` 的值看一下就很清楚了。**`middle` 不加 `1` 会造成死循环的代码**。

调试代码：

```go
package main

import "fmt"

func main() {
   fmt.Println(mySqrt(7))
}

func mySqrt(x int) int {
   // 特殊值判断
   if x == 0 || x == 1 {
      return x
   }
   left, right := 1, x/2
   for left < right {
      middle := (left + right) >> 1
      fmt.Println("left = ", left, ", right = ", right, ", middle = ", middle)
      if middle == x/middle { // 用除法防止溢出
         return middle
      } else if middle < x/middle { // 更新left（结果可能是middle）
         left = middle
      } else { // 更新right（结果不可能是right，因为都middle*middle大于x了）
         right = middle - 1
      }
   }
   return left
}
```

输出结果（死循环）：

```
left =  1 , right =  3 , middle =  2
left =  2 , right =  3 , middle =  2
left =  2 , right =  3 , middle =  2
left =  2 , right =  3 , middle =  2
left =  2 , right =  3 , middle =  2
left =  2 , right =  3 , middle =  2
...
```

**在区间只有 2 个数的时候**，本题 if、else 的逻辑区间的划分方式是：`[left, middle - 1]` 与         `[middle, right]`。**如果 middle下取整，在区间只有 2 个数的时候有 middle 的值等于 left，一旦进入分支 [middle, right] 区间不会再缩小，出现死循环。**

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220924201103.png)

**解决办法**：把取中间数的方式改成上取整。（即`middle := (left + right + 1) >> 1`）

> 这里的位运算其实就是`/2`操作。
>
> `middle := (left + right) >> 1`等同于`middle := (left + right) /2`
>
> 之所以使用位运算是因为位运算比除法更快。
>
> [轻松理解 left + ((right -left) ＞＞ 1)](https://blog.csdn.net/weixin_45935610/article/details/123121742)
>
> （其他人写`left + ((right -left) ＞＞ 1`（等同于`middle := left + (right - left) / 2`）是为了防止`left+right`的结果溢出了，但是我目前刷题的经验来go代码中并没有溢出）





