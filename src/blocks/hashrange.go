package blocks

import "fmt"

func HashRange() {
	hashMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Loop %d\n", i)
		for k, v := range hashMap {
			fmt.Println(k, v)
		}
	}

}
