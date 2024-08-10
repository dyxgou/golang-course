package types

import "fmt"

func Operations() {
	var x int = 10

	x *= 2 // x = x * 2

	fmt.Println("Go supports all the operations followed by and =")
	fmt.Println("+=, -=, /=, %=")
	fmt.Printf("x = %d", x)
}
