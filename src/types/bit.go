package types

import "fmt"

func BitManipulation() {
	b1 := 0b10011
	b2 := 0b11110

	fmt.Printf("b1 = %b, b2 = %b\n", b1, b2)

	fmt.Printf(
		"sum = %b, and = %b, or = %b, xor = %b, nand = %b\n",
		b1+b2,
		b1&b2,
		b1|b2,
		b1^b2,
		b1&^b2,
	)
}
