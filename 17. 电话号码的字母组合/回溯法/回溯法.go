package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// 数字到字母的映射
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := []string{}
	path := []byte{}
	backtrack(digits, 0, path, &res, phoneMap)
	return res
}

// backtrack 递归构建组合
// digits: 输入数字串
// index: 当前处理到第几个数字
// path: 当前已构建的字母组合（字节切片）
// res: 结果集指针
// phoneMap: 数字到字母的映射
func backtrack(digits string, index int, path []byte, res *[]string, phoneMap map[byte]string) {
	if index == len(digits) {
		// 得到一个完整组合，将其转换为字符串加入结果集
		*res = append(*res, string(path))
		return
	}
	// 获取当前数字对应的字母串
	letters := phoneMap[digits[index]]
	for i := 0; i < len(letters); i++ {
		// 做选择：将当前字母加入路径
		path = append(path, letters[i])
		// 递归处理下一个数字
		backtrack(digits, index+1, path, res, phoneMap)
		// 撤销选择：回溯
		path = path[:len(path)-1]
	}
}

func main() {

}
