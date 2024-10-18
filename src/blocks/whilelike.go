package blocks

import "fmt"

func WhileLikeFor() {
	name := "Diego"
	i := 0

	for name == "Diego" {
		if i > 10 {
			name = "Daniel"
		}

		fmt.Println(name)
		i++
	}
}
