package leetcode

func climbStairs(n int) int {
	a := make([]int, n+1)

	a[0], a[1] = 1, 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	return a[n]
}
