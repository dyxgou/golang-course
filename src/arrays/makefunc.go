package arrays

import "fmt"

func MakeFunc() {
	// The make func allows you to create types that you can specify the len and the capacity optionaly

	x := make([]int, 6)

	// This creates a slice with len 6

	fmt.Println(x)
}
