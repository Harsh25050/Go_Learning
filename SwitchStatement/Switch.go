package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")

	}

	// Duplicate cases are not allowed
	finger1 := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger1 {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 4: //duplicate case
	// 	fmt.Println("Another Ring")
	case 5:
		fmt.Println("Pinky")

	default: //default case
		fmt.Println("incorrect finger number")

	}

	// Multiple expressions in case
	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	// Expressionless switch
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}

	// Fallthrough
	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	switch num := 25; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}

}
