package combinatorics

func Plus(a int) int {
	return a + 1
}

func Multiply(a int) int {
	return a * 2
}

func Operations() map[int] func(int) int {
	return map[int] func(int) int{
		0: Plus,
		1:  Multiply,
	}
}

func Operate(a int, codeOperation int) int {
	return Operations()[codeOperation](a)
}