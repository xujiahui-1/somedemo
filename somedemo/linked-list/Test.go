package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val      int
	Children []*Node
}

func uniquePaths(m int, n int) int {
	//m宽  n长
	//dp数组的定义, ji:左边。 从起点到当前节点的走法
	dp := make([][]int, m)
	//初始化
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}

	}
	return dp[m-1][n-1]
}

func main() {
	uniquePaths(3, 7)
}
