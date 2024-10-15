package arrays

import "fmt"

// We we need a piece of data that is independent from the other, we can use copy to create a copy of the same data in a different direction of memory
func CopyFunc() {
	badSlice := []string{
		"This a slice",
		"Of strings",
		"Copied",
	}

	badCopy := badSlice[:2]
	badCopy[0] = "Corrupted data!"

	fmt.Println("bad copied array :", badCopy)

	goodCopy := make([]string, 2) // It takes the lenght and copy the elements till the slice is full

	n := copy(goodCopy, badSlice)
	goodCopy[0] = "Uncorrupted data"

	fmt.Println(n)
	fmt.Println(goodCopy)
}
