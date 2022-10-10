package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}

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
