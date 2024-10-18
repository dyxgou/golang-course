package structs

import "fmt"

type person struct {
	name string
	age  uint
	pet  string // It can be an interface
}

func Structures() {
	newPerson := person{
		name: "Diego",
		age:  18,
		pet:  "Nothing",
	}

	fmt.Println(newPerson)
}
