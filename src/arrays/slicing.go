package arrays

import "fmt"

func SlicingSlices() {
	// As in python, you can slice slices using :

	names := make([]string, 10)

	names[0] = "Diego"
	names[1] = "Daniel"
	names[2] = "Danna"
	names[3] = "Laura"
	names[4] = "Didier"
	names[5] = "Nubia"
	names[6] = "Nubia"
	names[7] = "Nubia"
	names[8] = "Nubia"
	names[9] = "Nubia"

	fmt.Println(names[:2])
}
