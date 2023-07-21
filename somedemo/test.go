package main

import "math"

func main() {
}

// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量,
// 最少数量 , 先取大数？
func numSquares(n int) int {

	//dp的定义，dp长度= n/2即可
	//dp[i]= i最少被几个数的平方能凑成i
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	//递推公式：dp[j] = min(dp[j - i * i] + 1, dp[j]);
	for i := 1; i <= n; i++ { //金币，
		// 遍历背包
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j-i*i]+1, dp[j])
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
