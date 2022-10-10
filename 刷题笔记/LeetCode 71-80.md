# 76. 最小覆盖子串

[LeetCode 76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)

## 滑动窗口

这题我并没有参考其他题解，是按照自己的想法做的。

按照滑动窗口的思想：

- 判断滑动窗口内的子串是否合法：因为字符串中既有大写字母又有小写字母，所以我就使用一个长度为58的数组来存储滑动窗口内的子串以及进行比较的字符串t（结合ASCII码，这种处理已经很常见了），对应check函数的写法如下：

  ```go
  func check(windowArr [58]int, tArr [58]int) bool {
     for i := 0; i < 58; i++ {
        if windowArr[i] < tArr[i] {
           return false
        }
     }
     return true
  }
  ```

- 当滑动窗口满足条件后，我们要开始收缩滑动窗口，直到滑动窗口不满足条件了，这里就得到了一个覆盖子串ans：

  ```go
  if check(windowArr, tArr) { // 如果符合条件了就要收缩滑动窗口以得到最小的覆盖子串
     for check(windowArr, tArr) {
        windowArr[s[left]-'A']--
        left++
     }
     ans = min(ans, s[left-1:i+1])
  }
  ```

- 维护一个`ans`，在遍历过程中可能得到多个符合条件的子串，我们要维护一个最小的覆盖子串，对应min函数的写法如下：

  ```go
  func min(ans string, s string) string {
     if ans == "" || len(s) < len(ans) {
        return s
     }
     return ans
  }
  ```

代码如下：

```go
func minWindow(s string, t string) string {
	ans := ""
	lengthS, lengthT := len(s), len(t)
	if lengthS < lengthT {
		return ""
	}
	// 得到t字符串对应的数组arrT（方便和滑动窗口类的字符串进行比较，判断滑动窗口是否合法）
	tArr := [58]int{}
	for _, c := range t {
		tArr[c-'A']++
	}
	// 准备维护滑动窗口
	left, i := 0, 0
	windowArr := [58]int{}
	for i < lengthS {
		windowArr[s[i]-'A']++
		if check(windowArr, tArr) { // 如果符合条件了就要收缩滑动窗口以得到最小的覆盖子串
			for check(windowArr, tArr) {
				windowArr[s[left]-'A']--
				left++
			}
			ans = min(ans, s[left-1:i+1])
		}
		i++
	}
	return ans
}

func check(windowArr [58]int, tArr [58]int) bool {
	for i := 0; i < 58; i++ {
		if windowArr[i] < tArr[i] {
			return false
		}
	}
	return true
}

func min(ans string, s string) string {
	if ans == "" || len(s) < len(ans) {
		return s
	}
	return ans
}
```

























