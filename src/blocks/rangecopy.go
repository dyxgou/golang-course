package blocks

import "fmt"

func RangeCopyValue() {
	array := [...]int{1, 2, 3, 4, 5, 65, 6, 7}

	for _, v := range array {
		v *= 2
		fmt.Println(v)
	}

	fmt.Println(array)
}
