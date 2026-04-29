package main

func timeRequiredToBuy(tickets []int, k int) int {
	// 初始化队列，存储所有人的索引
	queue := make([]int, len(tickets))
	for i := 0; i < len(tickets); i++ {
		queue[i] = i
	}
	time := 0
	// 剩余票数副本，避免修改原数组（也可直接操作原数组）
	remaining := make([]int, len(tickets))
	copy(remaining, tickets)
	for len(queue) > 0 {
		// 队首的人
		idx := queue[0]
		queue = queue[1:] // 出队
		// 买一张票
		remaining[idx]--
		time++
		// 如果买完最后一张票，且这个人正是 k，则返回时间
		if idx == k && remaining[idx] == 0 {
			return time
		}
		// 如果还有剩余票，重新入队
		if remaining[idx] > 0 {
			queue = append(queue, idx)
		}
	}
	return time // 理论上不会执行到这里，因为题目保证会完成
}

func main() {

}
