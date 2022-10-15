package main

func main() {

}

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

//func intersection(nums1 []int, nums2 []int) []int {
//	set1 := make(map[int]struct{}, 0)
//	set2 := make(map[int]struct{}, 0)
//	res := make([]int, 0)
//	// 两个切片的去重操作
//	for _, num := range nums1 {
//		if _, ok := set1[num]; !ok { // map中没有这个元素，添加
//			set1[num] = struct{}{}
//		}
//	}
//	tmp := make([]int, 0) // 去重后的nums2
//	for _, num := range nums2 {
//		if _, ok := set2[num]; !ok { // map中没有这个元素，添加
//			set2[num] = struct{}{}
//			tmp = append(tmp, num)
//		}
//	}
//	// 遍历去重后的nums2切片
//	for _, num := range tmp {
//		if _, ok := set1[num]; ok { // map中有这个元素，说明是两个切片的交集
//			res = append(res, num)
//		}
//	}
//	return res
//}
