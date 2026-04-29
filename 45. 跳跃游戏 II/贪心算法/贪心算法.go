package main

func jump(nums []int) int {
	// steps：当前已经跳跃的次数。
	// curEnd：当前步数所能到达的最远边界。
	// maxReach：在当前步数范围内，能够到达的最远位置。
	var steps, curEnd, maxReach int
	// ● 遍历数组（不包括最后一个元素），在遍历过程中：
	//  ○ 更新 maxReach = max(maxReach, i + nums[i])。
	//  ○ 当 i == curEnd 时，表示已经达到当前步数的边界，必须再跳一次，于是 steps++，并设置 curEnd = maxReach。
	//● 当 curEnd >= n-1 时，可以提前结束（但通常遍历完即可）。
	for i := 0; i < len(nums)-1; i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == curEnd {
			steps++
			curEnd = maxReach
		}
	}
	return steps
}

func main() {

}
