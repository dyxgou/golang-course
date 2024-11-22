package functions

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, len(vals))

	for i, v := range vals {
		out[i] = v + base
	}

	return out
}

func addToPointer(base int, vals []int) {
	// The composite types are pased by references

	for i, v := range vals {
		vals[i] = v + base
	}
}

func VariadicFuncs() {

	out := addTo(5, 1, 2, 3, 4, 5, 6, 67, 7)

	fmt.Println(out)
	vals := []int{
		1, 23, 4, 5, 6, 6, 76, 78,
	}

	addToPointer(12, vals)

	fmt.Printf("vals: %v\n", vals)
}
