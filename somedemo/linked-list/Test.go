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

// 有障碍物的不同走法
// dp数组的定义：i.j坐标点   dp[i][j]:走到该点的路径
// dp初始化； 最左一排和最上一排都是1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[i]))

	}
	//必须要不是障碍的时候我们才进行初始化1
	//一旦遇到障碍了，那么他后边的所有数都是走不通的，所以都赋值为0
	for i := 0; i < len(dp) && obstacleGrid[i][0] != 1; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < len(dp[0]) && obstacleGrid[0][i] != 1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)][len(obstacleGrid[0])]

}

func main() {
	test := [][]int{{0, 0, 0}, {0, 0, 0}, {1, 0, 0}}
	uniquePathsWithObstacles(test)
}
