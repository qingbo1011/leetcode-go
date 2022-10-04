package main

func main() {

}

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast = 0; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow+1] // 更新数组
	return slow + 1
}
