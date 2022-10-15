package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	// key为长度为26的int数组[26]int，value为string切片[]string
	mp := make(map[[26]int][]string) // 定义map
	for _, str := range strs {
		// 先将str转成长度为26的int数组
		tmp := [26]int{}
		for i := range str {
			tmp[str[i]-'a']++
		}
		// 然后将该数组和对应的str加入到map中（注意value是切片）
		mp[tmp] = append(mp[tmp], str)
	}
	res := make([][]string, 0, len(mp)) // len(map)即为map键值对数量（这里很显然res的cap是mp的键值对数量）
	for _, value := range mp {
		res = append(res, value)
	}
	return res
}
