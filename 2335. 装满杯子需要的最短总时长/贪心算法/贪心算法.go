package main

import "sort"

func main() {

}

func fillCups(amount []int) int {
	ans := 0
	for {
		sort.Ints(amount)
		if amount[1] == 0 { // 排序后，第二大的数已经为0了，说明只剩下一个数字需要处理，在最后直接ans + amount[2]即可
			break
		}
		amount[1]--
		amount[2]--
		ans++
	}
	return ans + amount[2]
}
