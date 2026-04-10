package main

func subarraySum(nums []int, k int) int {
	count := 0
	preSum := 0            // 当前前缀和，即nums[0]+...+nums[i]
	m := make(map[int]int) // 记录每个前缀和出现的次数
	m[0] = 1               // 空前缀（和为0，出现1次），用于处理从开头开始的子数组
	for i := 0; i < len(nums); i++ {
		preSum = preSum + nums[i]
		// 关键：查找之前有多少个前缀和等于preSum-k
		// 因为这些前缀和对应的位置到当前位置的子数组和恰好为k
		count = count + m[preSum-k]
		m[preSum]++ // 将当前前缀和存入哈希表，供后续元素使用
	}

	return count
}

func main() {

}
