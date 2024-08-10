package types

import "fmt"

func TypeConvertion() {
	var age int = 17
	var momAge int = 43
	var floatAge float64 = float64(age) + 0.7

	fmt.Printf("The ages added up are : %d\n", age+momAge)
	fmt.Printf("Float age %f, age type %T\n", floatAge, floatAge)
}
