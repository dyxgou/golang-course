package types

func Vars() {
	// The most explicit way to declare a variable in go is like this

	var name string = "Diego"

	// You can also declare a variable and leave it empty this is going to assign it its zero value
	var age int8

	var myAge, momAge int = 17, 43

	// If you have to declare multiple variables, you can wrap them into a declaration statement like so:

	var (
		name       string = "Diego"
		secondName string = "Alejandro"
		dadAge            = 42 // int
	)

	// Go has a way to declare variables that infers its own type as the default type of the value

	diegoAge := 17
}
