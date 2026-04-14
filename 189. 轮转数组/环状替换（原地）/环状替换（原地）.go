package main

func rotate(nums []int, k int) {
	k = k % len(nums) // 防止k过大溢出
	// 如果k被len(nums)整除说明反转后的结果就是它本身
	if k == 0 {
		return
	}
	count := 0 // 已移动的元素个数
	// 从start=0开始循环，直到 count == n(start++ 继续处理下一个环)
	for start := 0; count < len(nums); start++ {
		cur := start
		for {
			// 交换
			next := (cur + k) % len(nums)
			nums[next], nums[start] = nums[start], nums[next]
			cur = next
			count++
			if cur == start {
				break
			}
		}
	}
}

func main() {

}
