package arrays

import "fmt"

func Append() {
	names := []string{
		"Diego",
		"Daniel",
		"Heidy",
		"Laura",
		"Yaritza",
		"Santiago",
	}

	fmt.Printf("Names before = %v\n", names)

	names = append(names, "Nicolas") // you can appaned more than one value
	// names = append(names, "Nicolas", "Esteban", "Laura")

	fmt.Printf("Names after = %v\n", names)
}
