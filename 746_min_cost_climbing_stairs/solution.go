package main

import "math"

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make(map[int]int)
	return int(math.Min(float64(minCost(length-1, cost, dp)), float64(minCost(length-2, cost, dp))))
}

func minCost(i int, cost []int, dp map[int]int) int {
	if i < 0 {
		return 0
	}

	if i == 0 || i == 1 {
		return cost[i]
	}

	if _, ok := dp[i]; ok {
		return dp[i]
	}

	dp[i] = cost[i] + int(math.Min(float64(minCost(i-1, cost, dp)), float64(minCost(i-2, cost, dp))))
	return dp[i]
}
