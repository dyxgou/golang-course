package blocks

import "fmt"

func Shadowing() {
	x := 10

	if x > 5 {
		fmt.Printf("x: %v\n", x)

		x := 5

		fmt.Printf("x: %v\n", x)
	}

	fmt.Printf("x: %v\n", x)
}
