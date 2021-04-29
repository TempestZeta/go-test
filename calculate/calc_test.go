package calculate

import "fmt"

func ExampleConvertNumber() {
	numbers := []int{4, 2, 1}

	fmt.Println(ConvertNumber(numbers))
	// Output : 421
}
