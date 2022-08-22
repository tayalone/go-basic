package calculator

var sum = Sum

/*
Multiply some int
*/
func Multiply(a, b int) int {
	result := 0
	for i := 0; i < a; i++ {
		result += sum(b, 0)
	}
	return result
}
