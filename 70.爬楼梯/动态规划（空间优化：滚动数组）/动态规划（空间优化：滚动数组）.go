package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	// a 代表 dp[i-2], b 代表 dp[i-1]
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {

}
