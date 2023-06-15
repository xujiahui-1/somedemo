package main

func main() {
}

// n个节点的二叉搜索树的组合数=左0右2+左1右1+左2右0
// dp【i】 += dp【j】 * dp【i-j】，是根据搜索树的特性
func numTrees(n int) int {
	//dp定义 n节点数：dp[n]:n个节点有几种情况
	dp := make([]int, n+1)
	//初始化
	dp[0] = 1
	dp[1] = 1
	//递推公式 dp[3]=dp[0]*dp[2]+dp[1]*dp[1]+dp[2]*dp[0]
	//dp[n]=
	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
