package main

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func absInt(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x - 0
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
