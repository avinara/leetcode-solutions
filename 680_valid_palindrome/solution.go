package main

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return isSubStringPalindrome(s, i+1, j) || isSubStringPalindrome(s, i, j-1)
		}
		i++
		j--
	}
	return true
}

func isSubStringPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
