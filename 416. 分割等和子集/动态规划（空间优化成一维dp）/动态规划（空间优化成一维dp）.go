package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true // 空集可以凑出和为0
	for _, num := range nums {
		// 倒序遍历，确保每个数字只用一次
		for j := target; j >= num; j-- {
			if dp[j-num] {
				dp[j] = true
			}
		}
	}
	return dp[target]
}

func main() {

}
