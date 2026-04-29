package main

func timeRequiredToBuy(tickets []int, k int) int {
	// 对于位置 k 的人，他需要买 tickets[k] 张票。在他购买的过程中，其他人也会同步购票。
	//  关键在于：在他完成最后一票之前，每个人能买的票数受限于他自身的需求量以及 tickets[k] 的轮次。
	//● 对于 k 左边 的人（索引 < k）：他们能买到的票数最多是 min(tickets[i], tickets[k])，
	//  因为当 k 买完 tickets[k] 张时，左边的人最多也被服务了 tickets[k] 次（如果他们的需求更多，则会被截断）。
	//● 对于 k 自己：需要 tickets[k] 秒。
	//● 对于 k 右边 的人（索引 > k）：他们能买到的票数最多是 min(tickets[i], tickets[k]-1)，
	//  因为 k 完成最后一次购票的那一轮，右边的人没有机会再买（他们排在 k 后面，而 k 买到最后一票后就离开了，那轮不会轮到他们）。
	total := 0
	for i := 0; i < len(tickets); i++ {
		if i == k {
			total += tickets[k]
		} else if i < k {
			total += min(tickets[i], tickets[k])
		} else {
			total += min(tickets[i], tickets[k]-1)
		}
	}
	return total
}

func main() {

}
