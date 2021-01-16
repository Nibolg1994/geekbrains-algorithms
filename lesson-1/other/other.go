package other

func Max(a int64, b int64, c int64) int64 {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}
