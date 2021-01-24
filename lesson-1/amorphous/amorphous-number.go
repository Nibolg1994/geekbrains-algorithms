package amorphous

import (
	"strconv"
)

func IsAmorphous(number int64) bool {
	squareStr := strconv.FormatInt(number*number, 10)
	numberStr := strconv.FormatInt(number, 10)
	n := len(numberStr)
	squareN := len(squareStr)
	for i := squareN - n; i < squareN; i++ {
		if squareStr[i] != numberStr[i-squareN+n] {
			return false
		}
	}

	return true
}

func GenerateNumbers(N int64) []int64 {
	var slice []int64 = make([]int64, 0, N)
	var i int64
	for i = 1; i <= N; i++ {
		if IsAmorphous(i) {
			slice = append(slice, i)
		}
	}
	return slice
}
