package blocks

import (
	"fmt"
	"math/rand"
)

func IfStatement() {
	n := rand.Intn(10)
	fmt.Printf("n: %v\n", n)

	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}
