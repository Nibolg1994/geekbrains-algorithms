package combinatorics

func FindSolution(start int, end int) int {
	if start > end {
		return 0
	}
	if start == end {
		return 1
	}
	return FindSolution(Operate(start, 1), end) + FindSolution(Operate(start, 0), end)
}