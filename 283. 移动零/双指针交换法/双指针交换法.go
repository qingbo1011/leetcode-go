package main

func moveZeroes(nums []int) {
	// left 指向当前已处理好的非零元素的下一个位置（即下一个非零元素应该放置的位置）
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

func main() {
}
