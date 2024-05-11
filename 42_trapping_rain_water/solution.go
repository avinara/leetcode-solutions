package main

import "math"

func trap(height []int) int {
	totalWaterTrapped := 0
	for i := 0; i < len(height); i++ {
		maxLeft := findMax(height[:i])
		maxRight := findMax(height[i:])
		if maxLeft < 0 || maxRight < 0 {
			continue
		}
		waterTrapped := int(math.Min(float64(maxLeft), float64(maxRight))) - height[i]
		if waterTrapped < 0 {
			continue
		}
		totalWaterTrapped += waterTrapped
	}
	return totalWaterTrapped
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
