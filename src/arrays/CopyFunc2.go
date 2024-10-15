package arrays

import "fmt"

func CopyFunc2() {
	slice := []int{12, 3, 4, 5}
	array := [4]int{5, 7, 8, 9}

	copySlice := make([]int, 2)

	copy(copySlice, array[:])

	fmt.Println(copySlice)

	copy(array[:], slice)

	fmt.Println("Array", array)
}
