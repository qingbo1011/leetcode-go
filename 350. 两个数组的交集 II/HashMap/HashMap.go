package main

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	mp := make(map[int]int, 0) // key为nums1中的数字，value为该数字在nums1出现的次数
	for _, num := range nums1 {
		mp[num]++
	}
	for _, num := range nums2 {
		if mp[num] > 0 { // num既存在于mp中，且value>0（这里不用value, ok := mp[num]; ok && value > 0，
			// 因为go中如果map的key不存在，mp[key]会得到value的0值 ）
			res = append(res, num)
			mp[num]-- // 别忘了value-1
		}
	}
	return res
}
