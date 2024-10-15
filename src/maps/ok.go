package maps

import (
	"fmt"
	"log"
)

func Ok() {
	wins := map[string]int{
		"Diego":  3,
		"Daniel": 4,
		"Laura":  5,
	}

	v, ok := wins["Alejandro"]

	if !ok {
		log.Fatal("That key doesnt exists")
	}

	fmt.Printf("Wins = %d\n", v)
}
