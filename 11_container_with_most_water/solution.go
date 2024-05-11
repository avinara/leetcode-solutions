package main

import "math"

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	i := 0
	j := len(height) - 1
	maxArea := 0
	for i < j {
		minHeight := int(math.Min(float64(height[i]), float64(height[j])))
		width := j - i
		area := width * minHeight
		if area > maxArea {
			maxArea = area
		}

		if minHeight == height[i] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
