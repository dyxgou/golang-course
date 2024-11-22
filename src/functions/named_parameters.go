package functions

import "fmt"

type myFuncParams struct {
	FirstName  string
	SecondName string
	age        uint
}

func myFunc(opts myFuncParams) {
	fmt.Println(opts)
}

func NamedParameters() {
	myFunc(myFuncParams{
		SecondName: "Alejandro",
		age:        18,
	})

	myFunc(myFuncParams{
		FirstName: "Diego",
	})

}
