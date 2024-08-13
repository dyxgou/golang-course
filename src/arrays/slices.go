package arrays

import "fmt"

func Slices() {
	// One of the main things about slice is that we dont specify the length of the slice

	// [...]int -> Makes an array
	// []int -> Makes a Slice

	ints := []int{1, 5, 6, 7, 8}

	fmt.Printf("slices = %v", ints)

	ints2 := []int{5, 1, 2, 4, 6}

	fmt.Printf("slices = %v\n", ints2)

	// The slices are not comparable this means that if you try to compare two slices it is going to trow a compile error

	// Invalid operation: Slices can only be comparable to nil

	// if ints == ints2 {
	// 	// Do something
	// }

	// In go `nil` is the identifier to represent the lack of a value which is different from the `null` seen in another languages

	fmt.Printf("is Slice equal %t\n", isSliceEqual(ints, ints))

	// The len func returns 0 if you pass it a empty slice (a nil slice)

	var empty []int

	fmt.Printf("empty size = %d\n", len(empty))
}

func isSliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		n1 := s1[i]
		n2 := s2[i]

		if n1 != n2 {
			return false
		}
	}

	return true
}
