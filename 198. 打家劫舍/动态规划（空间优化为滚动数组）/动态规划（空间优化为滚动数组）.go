package main

func rob(nums []int) int {
	prev2, prev1 := 0, 0
	for _, num := range nums {
		cur := max(prev1, prev2+num)
		prev2, prev1 = prev1, cur
	}

	return prev1
}

func main() {

}
