package dp

import "testing"

func TestLT62(t *testing.T) {
	t.Log(uniquePaths(3, 2))
}
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}
