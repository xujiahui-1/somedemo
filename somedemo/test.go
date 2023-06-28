package main

import "math"

func main() {

}

//加加减减变成target有几种方法

func findTargetSumWays(nums []int, target int) int {
	// jia+jin=sum
	//jia-jin=target
	//jia=target+jin
	//sum=target+jin+jin= target+jin*2
	//jin=(sum-target)/2

	sum := 0
	for _, i := range nums {
		sum += i
	}
	if (sum-target)%2 == 1 {
		return 0
	}
	if abs(target) > sum {
		return 0
	}
	bagelength := (sum - target) / 2
	dp := make([]int, bagelength+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagelength; j >= nums[i]; j++ {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagelength]
}
func abs(x int) int {
	return int(math.Abs(float64(x)))
}
