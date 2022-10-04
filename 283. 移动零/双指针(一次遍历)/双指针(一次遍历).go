package main

func main() {

}

func moveZeroes(nums []int) {
	slow, fast := 0, len(nums)-1
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
	}
}
