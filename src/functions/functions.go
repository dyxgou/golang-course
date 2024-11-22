package functions

import (
	"errors"
	"fmt"
	"log"
)

func division(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("Invalid operation")
	}

	return numerator / denominator, nil
}

func Functions() {
	result, err := division(5, 8)

	if err != nil {
		log.Fatal("Cant divide that shit", err)
	}

	fmt.Printf("result: %v\n", result)
}
