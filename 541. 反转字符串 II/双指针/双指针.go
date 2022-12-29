package main

func main() {

}

func reverseStr(s string, k int) string {
	strBytes := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		if i+k < len(s) { // 反转字符中的前k个字符
			reverseString(strBytes[i : i+k])
		} else { // 如果剩余字符少于k个，则将剩余字符全部反转
			reverseString(strBytes[i:len(s)])
		}
	}
	return string(strBytes)
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
