package main

func main() {

}

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, num := range nums {
		if sum < 0 {
			sum = num
		} else {
			sum = sum + num
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
