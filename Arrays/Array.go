package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// Declaration
	var a [3]int //int array with length 3
	fmt.Println(a)
	a1 := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a1)

	// Arrays are value types

	a2 := [...]string{"USA", "China", "India", "Germany", "France"}
	b2 := a2 // a copy of a is assigned to b
	b2[0] = "Korea"
	fmt.Println("a is ", a2)
	fmt.Println("b is ", b2)
	// Length of an array
	fmt.Println("length of a is", len(a2))

	// Iterating arrays using range
	a3 := [...]float64{67.7, 103.8, 67.098, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a3[i])
	}

	// Multidimensional arrays
	a4 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a4)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}
