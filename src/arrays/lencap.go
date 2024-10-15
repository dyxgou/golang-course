package arrays

import "fmt"

func LenCap() {
	names := make([]string, 10)

	// All the arrays in golang that we create with a certatin len they also have a capacity which if the len is less than 1024, they grow by 2x

	fmt.Println(names)
	fmt.Printf("names len = %d\n", len(names))
	fmt.Printf("names cap= %d\n", cap(names))

	for i := 0; i < len(names); i++ {
		names[i] = "Diego"
	}

	names = append(names, "Daniel")

	fmt.Printf("names len = %d\n", len(names))
	fmt.Printf("names cap= %d\n", cap(names))

	fmt.Println(names)
}
