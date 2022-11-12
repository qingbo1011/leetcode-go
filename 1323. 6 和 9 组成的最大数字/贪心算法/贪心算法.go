package main

import "strconv"

func main() {

}

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	strByte := []byte(str)
	for i := 0; i < len(strByte); i++ {
		if strByte[i] == '6' {
			strByte[i] = '9'
			break
		}
	}
	ans, _ := strconv.Atoi(string(strByte))
	return ans
}
