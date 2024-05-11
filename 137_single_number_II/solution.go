package main

func singleNumber(nums []int) int {

	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; !ok {
			numMap[nums[i]] = 1
		} else {
			numMap[nums[i]]++
		}
	}

	var singleNum int
	for i := 0; i < len(nums); i++ {
		if numMap[nums[i]] == 1 {
			singleNum = nums[i]
		}
	}
	return singleNum
}
