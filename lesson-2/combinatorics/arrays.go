package combinatorics

import (
	"math"
	"strconv"
)


func FindSolutionWithArray(start int, end int) int {
	i := start
	count := 0
	steps := end - start
	var Empty struct{}
	results := make(map[string] interface{})
	for i = 0; i < int(math.Pow(2, float64(steps))); i++ {
		result := ""
		sum := start
		for j := 0; j < steps; j++ {
			codeOperator := 0
			if i & (1 << j) != 0 {
				codeOperator = 1
			}
			sum = Operate(sum, codeOperator)
			result += strconv.Itoa(codeOperator)
			if sum >= end {
				break
			}

		}
		if sum == end {
			if _, ok := results[result]; !ok {
				count++
				results[result] = Empty
			}
		}
	}
	return count
}


