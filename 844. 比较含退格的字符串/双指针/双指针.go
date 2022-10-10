package main

func main() {

}

func backspaceCompare(s string, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 { // 大循环，在里面遍历S和T
		// 找到s需要比较的字符
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if s[i] != '#' && skipS != 0 {
				skipS--
				i--
			} else { // 找到s需要比较的字符了
				break
			}
		}
		// 与t中需要比较的字符进行对比
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if t[j] != '#' && skipT != 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		// 开始比较两个字符
		if i >= 0 && j >= 0 { // 如果index=0的位置上为 '#'，则 i,j会为-1，随意要判断一下
			if s[i] != t[j] {
				return false
			}
		} else if !(i < 0 && j < 0) { // 如果i和j不是都遍历结束了，说明两个字符串中其中一个有待比较的字符而另一个已经遍历结束了，返回false
			return false
		}
		i--
		j--
	}
	return true
}
