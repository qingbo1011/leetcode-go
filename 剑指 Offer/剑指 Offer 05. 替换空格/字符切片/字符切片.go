package main

func main() {

}

func replaceSpace(s string) string {
	res := make([]byte, 0)
	strBytes := []byte(s)
	for i := 0; i < len(s); i++ {
		if strBytes[i] == ' ' { // 是空格
			res = append(res, '%', '2', '0')
		} else { // 不是空格直接append该字符即可
			res = append(res, strBytes[i])
		}
	}
	return string(res)
}
