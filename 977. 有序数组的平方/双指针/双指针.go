package main

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	println(sortedSquares(nums))

}

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums)) // 创建一个len为len(nums)的答案切片
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if v, k := nums[left]*nums[left], nums[right]*nums[right]; v < k { // 右边的平方大
			ans[i] = k
			right--
		} else {
			ans[i] = v
			left++
		}
	}
	return ans
}
