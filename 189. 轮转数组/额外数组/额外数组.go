package main

func rotate(nums []int, k int) {
	k = k % len(nums) // 防止k过大溢出
	// 如果k被len(nums)整除说明反转后的结果就是它本身
	if k == 0 {
		return
	}
	tmp := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, tmp)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

}
