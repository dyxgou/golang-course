package arrays

import "fmt"

// This happends because of when we slice an array we dont copy it, instead the two variables share the same memory thus making it able to overwrite the adjacent element
func ConfusingSlices() {
	x := make([]int, 0, 5)

	x = append(x, 1, 2, 3, 4) // We append four items

	y := x[:2:2] // Slice form 0 to 2 [0 , 2)
	z := x[2:]   // Slice from [2, n]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50) // [1,2,30,40,50]

	x = append(x, 60) // [1,2,30,40,60]
	z = append(z, 70) //[1,2,30,40,70]

	fmt.Println("Original Slice :", x)
	fmt.Println(y)
	fmt.Println(z)

	y[0] = 123123 // To avoid colitions between slices that share memory we can use the thirt part slice expresion which tells how much memory they can share

	// This have the same effect on arrays
}
