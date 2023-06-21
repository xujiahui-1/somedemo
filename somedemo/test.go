package main

func main() {

}

// stones[i] 表示第 i 块石头的重量
// 把石头分成尽量相等的两堆石头
func lastStoneWeightII(stones []int) int {
	//dp[j]=装满容量为j的背包的最大重量
	//dp[j]=maxx(dp[j], dp[j-stones[i]]+stones[i])

	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2 //小的
	//dp
	dp := make([]int, target+2)

	//遍历
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = maxx(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - 2*dp[target]

}
func maxx(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
