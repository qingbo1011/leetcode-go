package main

func main() {

}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	mp := make(map[int]int, 0)
	n := len(nums1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mp[nums1[i]+nums2[j]]++
		}
	}
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			//if _, ok := mp[0-(nums3[k]+nums4[l])]; ok { // 说明nums1[i]+nums2[j]+nums3[k]+nums4[l]==0，即我们想要的结果
			//	ans = ans + mp[0-(nums3[k]+nums4[l])]
			//}
			ans = ans + mp[0-(nums3[k]+nums4[l])] // 可以直接这样写，因为go中如果mp[key]中的key不存在，取出来的value是0
		}
	}
	return ans
}
