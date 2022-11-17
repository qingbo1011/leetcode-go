# 344. 反转字符串

[LeetCode 344. 反转字符串](https://leetcode.cn/problems/reverse-string/)

## 双指针

直接使用双指针实现反转。思路很简单不用多说。

代码如下：

```go
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
```

















# 349. 两个数组的交集

[LeetCode 349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/)

## HashSet

之前在项目中已经写过切片之间的交集并集差集了，思路不用多说。

求差集的话，HashSet就够了。这里提一下go中的set。go没有set这个数据结构，但是我们可以通过定义map并指定value为，空结构体`struct{}`来实现set。因为在Go中，**空结构体`struct{}`是不占任何内存的**。

> 在Go中，空结构体`struct{}`是不占任何内存的。因为空结构体不占据内存空间，因此被广泛作为各种场景下的占位符使用。一是节省资源。二是空结构体本身就具备很强的语义，即这里不需要任何值，仅作为占位符。（其中经典的用法就是用来实现集合Set。这里可以引申讲讲[Go空结构体struct{}](https://www.xiuxiubiu.com/archives/go%E7%A9%BA%E7%BB%93%E6%9E%84%E4%BD%93struct)）
>

代码如下：

```go
func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{}, 0)
	set2 := make(map[int]struct{}, 0)
	res := make([]int, 0)
	// 两个切片的去重操作
	for _, num := range nums1 {
		if _, ok := set1[num]; !ok { // map中没有这个元素，添加
			set1[num] = struct{}{}
		}
	}
	tmp := make([]int, 0) // 去重后的nums2
	for _, num := range nums2 {
		if _, ok := set2[num]; !ok { // map中没有这个元素，添加
			set2[num] = struct{}{}
			tmp = append(tmp, num)
		}
	}
	// 遍历去重后的nums2切片
	for _, num := range tmp {
		if _, ok := set1[num]; ok { // map中有这个元素，说明是两个切片的交集
			res = append(res, num)
		}
	}
	return res
}
```

其实感觉这样写代码还是有点臃肿的。可以使用一个set来实现，注意在遍历nums2时要`delete`键值对即可。

具体代码如下：

```go
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	res := make([]int, 0)
	for _, num := range nums1 { // 遍历nums1获取set
		if _, ok := set[num]; !ok { // 元素不存在map中，添加
			set[num] = struct{}{}
		}
	}
	for _, num := range nums2 { // 遍历nums2，根据上面得到的set就能获取交集
		if _, ok := set[num]; ok { // 元素存在map中，说明是交集的元素
			res = append(res, num)
			delete(set, num) // 删除这个键值对，避免交集res中出现重复元素
		}
	}
	return res
}
```

# 350. 两个数组的交集 II

[LeetCode 350. 两个数组的交集 II](https://leetcode.cn/problems/intersection-of-two-arrays-ii/)

## HashMap

稍微分析一下就有思路了：

- nums1使用HashMap存储，**key为nums1中的数字**，**value为该数字在nums1出现的次数**
- 遍历nums2，当遍历到的num存在map中，且value>0时，将该数字num添加到结果中，并`mp[num]--`

思路很简单，一下就能搞懂。

代码如下：

```go
func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	mp := make(map[int]int, 0) // key为nums1中的数字，value为该数字在nums1出现的次数
	for _, num := range nums1 {
		mp[num]++
	}
	for _, num := range nums2 {
		if mp[num] > 0 { // num既存在于mp中，且value>0（这里不用value, ok := mp[num]; ok && value > 0，因为go中如果map的key不存在，mp[key]会得到value的0值）
			res = append(res, num)
			mp[num]-- // 别忘了value-1
		}
	}
	return res
}
```















