package blocks

import (
	"fmt"
	"math/rand"
)

func ShadowingStatements() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// fmt.Println(n) This returns an error of undefined : n
}
