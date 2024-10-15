package maps

import "fmt"

func Maps() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	totalWins["Kittens"]++
	totalWins["Lions"] = 3

	for k, v := range totalWins {
		fmt.Printf("key = %s, value %d\n", k, v)
	}
}
