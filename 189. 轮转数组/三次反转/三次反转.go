package main

func rotate(nums []int, k int) {
	k = k % len(nums) // 防止k过大溢出
	// 如果k被len(nums)整除说明反转后的结果就是它本身
	if k == 0 {
		return
	}
	// 通过反转整个数组，然后反转前 k 个元素，再反转后 n-k 个元素，即可实现右旋 k 步
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {

}
