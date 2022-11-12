# 121. 买卖股票的最佳时机

[LeetCode 121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)

## 单调栈

这题看到之后，首先想到的就是使用单调栈，**保证栈顶永远大于栈底**。

具体思路用文字不好说，这里使用单调栈的原因跟接雨水类似，要找到**低-高-低**结构。草稿图如下：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220502151353.png)

代码如下：

```go
func maxProfit(prices []int) int {
	stack := make([]int, 0)
	ans := 0
	for _, price := range prices {
		for len(stack) > 0 && price < stack[len(stack)-1] {
			if len(stack) >= 2 { // 保证能出现低-高-低结构
				ans = max(ans, stack[len(stack)-1]-stack[0])
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, price)
	}
	if len(stack) >= 2 { // 遍历结束后，栈内还有2个或以上元素
		ans = max(ans, stack[len(stack)-1]-stack[0])
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
```

## 一次遍历

上面单调栈的写法，是之前用Java刷题是自己想出来的。可能不太好理解，效率也不高。

[官方题解](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/solution/121-mai-mai-gu-piao-de-zui-jia-shi-ji-by-leetcode-/)

假设给定的数组为：`[7, 1, 5, 3, 6, 4]`，在图表上绘制给定数组中的数字，我们将会得到：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20221030213249.png)

根据题意我们很容易分析出：**我们要在历史最低价中买入股票，并在其之后的最高值卖出股票**。

用`minPrice`记录一个历史最低价格 ，在那之后第`i`天卖出股票能得到的利润就是`prices[i]-minPrice`。

使用`maxProfit`来记录最大利润只需要遍历价格数组一遍，记录历史最低点，然后在后面的遍历中：

- 如果这一天的股价`prices[i]`比我们目前记录的最低价`minPrice`还要低（`prices[i] < minPrice`），那么我们更新历史最低价：`minPrice = prices[i]`
- 如果在这一天卖出股票得到的利润`prices[i] - minPrice`比当前记录的最大利润`maxProfit`还要打（`prices[i]-minPrice > maxProfit`），那么我们更新最大利润：`maxProfit = prices[i] - minPrice`

代码如下：

```go
func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
```







