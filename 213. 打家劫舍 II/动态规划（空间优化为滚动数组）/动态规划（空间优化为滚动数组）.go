package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	case1 := robRange(nums[1:])           // 偷最后一家,则不能偷第一家
	case2 := robRange(nums[:len(nums)-1]) // 偷第一家,则不能偷最后一家
	return max(case1, case2)
}

func robRange(nums []int) int {
	prev2, prev1 := 0, 0
	for _, num := range nums {
		cur := max(prev1, prev2+num)
		prev2, prev1 = prev1, cur
	}

	return prev1
}

func main() {

}
