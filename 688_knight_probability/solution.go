package main

func knightProbability(n int, k int, row int, column int) float64 {

	directions := [][]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}

	if row < 0 || row >= n || column < 0 || column >= n {
		return 0
	}
	if k == 0 {
		return 1
	}

	var res float64
	res = 0
	for i := 0; i < len(directions); i++ {
		dir := directions[i]
		res += float64(knightProbability(n, k-1, row+dir[0], column+dir[1]) / 8)
	}
	return res
}
