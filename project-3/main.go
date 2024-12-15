package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World!")

	var result, remainder, err = intDiv(5, 1)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v with no remainder\n", result)
	} else {
		fmt.Printf("The result of the division is %v with a remainder of %v\n", result, remainder)
	}

	switch {
	case err != nil:
		fmt.Printf("%s\n", err.Error())
	case remainder == 0:
		fmt.Printf("The result of the division is %v with no remainder\n", result)
	default:
		fmt.Printf("The result of the division is %v with a remainder of %v\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close at all")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDiv(numerator int, denominator int) (int, int, error) {
	var err error // default type is nil

	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
