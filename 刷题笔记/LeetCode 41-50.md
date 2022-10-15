# 45. 跳跃游戏 II

[LeetCode 45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)

## 贪心算法







# 49. 字母异位词分组

[LeetCode 49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/)

## 数组

这题一开始我的想法是：

- 使用数组+ASCII码的形式来判断两个字符串是否是字母异位词（Go中数组是可以直接比较的，这个思路就可以直接参考[LeetCode 242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/)了）
- 定义一个exists数组，来存储在strs中已经放到res中的索引（以后遍历就不需要再考虑该索引了）

代码可能看起来比较麻烦，但是思路还是很好懂的。

代码如下：

```go
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	exist := make([]int, 0) // 存放在strs中已经放到res中的索引
	for i := 0; i < len(strs); i++ {
		if !isExist(i, exist) { // 该字符串并没有被放到res中
			exist = append(exist, i) // 先添加到exist中
			temp := []string{strs[i]}
			for j := i + 1; j < len(strs); j++ {
				if isAnagram(strs[i], strs[j]) {
					exist = append(exist, j) // 先添加到exist中
					temp = append(temp, strs[j])
				}
			}
			res = append(res, temp)
		}
	}
	return res
}

// 判断两个字符串是否为字母异位词
func isAnagram(s string, t string) bool {
	lengthS, lengthT := len(s), len(t)
	if lengthS != lengthT {
		return false
	}
	sArr, tArr := [26]int{}, [26]int{}
	for i := 0; i < lengthS; i++ {
		sArr[s[i]-'a']++
		tArr[t[i]-'a']++
	}
	return sArr == tArr // go中，数组是可以直接比较的
}

// 判断i是否在exist切片中
func isExist(i int, exist []int) bool {
	for _, ex := range exist {
		if i == ex {
			return true
		}
	}
	return false
}
```

## 哈希表+数组

其实上面靠纯数组的处理方式是比较偏激的，不使用哈希表，空间换时间。

在看了[官方题解](https://leetcode.cn/problems/group-anagrams/solution/zi-mu-yi-wei-ci-fen-zu-by-leetcode-solut-gyoc/)后，有了另外的思路。使用`map`，其中：

- **key**为长度为26的int数组`[26]int`
- **value**为string切片`[]string`

map这样设置之后，思路就很明确了，直接看代码即可。

代码如下：

```go
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
```

> 复杂度分析：
>
> - 时间复杂度：O(n(k+∣Σ∣))，其中 n 是strs 中的字符串的数量，k 是 strs 中的字符串的的最大长度，Σ 是字符集，在本题中字符集为所有小写字母，∣Σ∣=26。需要遍历 n 个字符串，对于每个字符串，需要 O(k) 的时间计算每个字母出现的次数，O(∣Σ∣) 的时间生成哈希表的键，以及O(1) 的时间更新哈希表，因此总时间复杂度是O(n(k+∣Σ∣))。
> - 空间复杂度O(n(k+∣Σ∣))，其中 n 是strs 中的字符串的数量，k 是 strs 中的字符串的最大长度，Σ 是字符集，在本题中字符集为所有小写字母，∣Σ∣=26。需要用哈希表存储全部字符串，而记录每个字符串中每个字母出现次数的数组需要的空间为 O(∣Σ∣)，在渐进意义下小于 O(n(k+∣Σ∣))，可以忽略不计。







