package main

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}
