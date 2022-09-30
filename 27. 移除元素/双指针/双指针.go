package main

func main() {

}

func removeElement(nums []int, val int) int {
	length := len(nums)
	slow, fast := 0, 0
	for fast = 0; fast < length; fast++ {
		if nums[fast] != val { // 是新数组成员，要更新慢指针
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow] // 更新一下数组
	return slow
}
