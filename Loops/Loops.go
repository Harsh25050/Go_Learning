package main

import (
	"fmt"
)

func main() {
	// For Loop
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d",i)
	}
// Break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	// continue
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}


	// nested
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}


	// Labels
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
		}

	}


	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	// Infinite Loop
	for {
		fmt.Println("Hello World")
	}
}