package main

func longestCommonPrefix(strs []string) string {
	smallString := strs[0]
	length := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if length > len(strs[i]) {
			length = len(strs[i])
			smallString = strs[i]
		}
	}

	var prefix string
	for i := 0; i < length; i++ {
		count := 0
		for _, str := range strs {
			if str[i] == smallString[i] {
				count++
			}
		}
		if count == len(strs) {
			prefix = prefix + string(smallString[i])
		} else {
			break
		}
	}
	return prefix
}
