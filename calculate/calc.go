package calculate

import "math"

func CalculateNumber(question string) int {
	var result int = 0

	return result
}

func ConvertNumber(numbers []int) int {
	size := len(numbers)

	var result int = 0

	for i := 0; i < size; i++ {
		result += numbers[i] * int(math.Pow(10.0, (float64(size-1-i))))
	}

	return result
}

func add(a int, b int) int {
	return a + b
}
