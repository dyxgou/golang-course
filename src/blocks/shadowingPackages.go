package blocks

import "fmt"

func ShadowingPackages() {
	// When you are using shadowing variables, you gotta be careful you dont shadow a package

	fmt.Println("Shadowed packaged")

	// fmt := "Hola"
	//
	// fmt.Println("Shadowed package")
}
