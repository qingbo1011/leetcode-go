package main

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0) // 长度为k的单调队列（从左到右递减）注意存储的是nums索引！
	for i := 0; i < len(nums); i++ {
		// 1.移除队尾所有小于等于当前值的元素（维护单调递减）
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		// 2.当前索引入队
		deque = append(deque, i)
		// 3.移除队首已经离开窗口的下标
		left := i - k + 1
		if deque[0] < left {
			deque = deque[1:]
		}
		// 4.当窗口大小达到k时，记录结果(等i遍历到往左能有窗口的长度后，就会一直走这段逻辑了)
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}

func main() {

}
