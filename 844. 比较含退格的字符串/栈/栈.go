package main

func main() {
	backspaceCompare("y#fo##f", "y#f#o##f")
}

func backspaceCompare(s string, t string) bool {
	stackS := make([]byte, 0)
	stackT := make([]byte, 0)
	for i := range s {
		if len(stackS) != 0 && s[i] == '#' { // 栈不为空且遇到删除符号，栈顶元素出栈
			stackS = stackS[:len(stackS)-1]
		} else if s[i] != '#' { // 遇到非'#'，入栈
			stackS = append(stackS, s[i])
		}
	}
	for i := range t {
		if len(stackT) != 0 && t[i] == '#' { // 栈不为空且遇到删除符号，栈顶元素出栈
			stackT = stackT[:len(stackT)-1]
		} else if t[i] != '#' { // 遇到非'#'，入栈
			stackT = append(stackT, t[i])
		}
	}
	// 开始比较两个栈是否相同
	if len(stackS) != len(stackT) {
		return false
	}
	for i := 0; i < len(stackS); i++ {
		if stackS[i] != stackT[i] {
			return false
		}
	}

	return true
}
