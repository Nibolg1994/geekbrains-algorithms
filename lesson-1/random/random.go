package random

import "time"

func GenerateNumber(N int64) int64 {
	return time.Now().UnixNano() % (N + 1)
}
