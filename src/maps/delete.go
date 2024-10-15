package maps

import (
	"fmt"
	"log"
)

func Delete() {
	wins := map[string]int{
		"Diego":  3,
		"Daniel": 4,
		"Laura":  5,
	}

	key := "Diego"
	v, ok := wins[key]

	if !ok {
		log.Fatal("That key doesnt exists")
	}

	fmt.Printf("v: %v\n", v)
	delete(wins, key)
	fmt.Printf("wins: %v\n", wins)
}
