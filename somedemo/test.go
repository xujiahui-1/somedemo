package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"10000", "11000"}
	findMaxForm(str, 5, 5)
}

// m =0   n =1
func findMaxForm(strs []string, m int, n int) int {
	//dp数组的定义： dp[m][n]代表:m个0，n个1 能存放多少个数字
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	//初始化
	dp[0][0] = 0
	//递推公式
	for i := 0; i < len(strs); i++ { //遍历数
		//计算当前数子有几个0 几个1
		num0 := strings.Count(strs[i], "0")
		num1 := strings.Count(strs[i], "1")

		//遍历背包    m =0   n =1
		for j := m; j > 0; j-- {
			for k := n; k > 0; k-- {
				if j >= num0 && k >= num1 { //说明当前背包可以放在当前数字
					dp[j][k] = max(dp[j][k], dp[j-num0][k-num1]+1)
				}

			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
