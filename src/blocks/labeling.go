package blocks

import "fmt"

func Labeling() {
outer:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, " is even")

		case i%3 == 0:
			fmt.Println(i, " is divisible by 3 and not for 2")

		case i%7 == 0:
			fmt.Println(i, " exit the loop")
			break outer
		default:
			fmt.Println(i, "is boring")
		}
	}
}
