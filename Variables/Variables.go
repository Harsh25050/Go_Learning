package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring a single variable
	var age int // // variable declaration
	fmt.Println("My age is", age)
	age = 29 //assignment
	fmt.Println("My age is", age)
	age = 54 //assignment
	fmt.Println("My new age is", age)

	// Declaring a variable with an initial value
	var age2 int = 24 // // variable declaration with initial value
	fmt.Println("My age is", age2)

	// Type inference
	var age3 = 29 // type will be inferred
	fmt.Println("My age is", age3)

	// Multiple variable declaration
	var width, height int = 100, 50 //declaring multiple variables

	fmt.Println("width is", width, "height is", height)
	var width1, height1 = 200, 66 //"int" is dropped

	fmt.Println("width is", width1, "height is", height1)

	var (
		name    = "Harsh"
		age4    = 21
		height3 int
	)
	fmt.Println("my name is", name, ", age is", age4, "and height is", height3)

	// Short hand declaration
	name, age5 := "Harsh", 21 //short hand declaration

	fmt.Println("my name is", name, "age is", age5)

	// name, age := "naveen" //error

	// fmt.Println("my name is", name, "age is", age)

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	a1, b1 := 145.8, 543.8
	c1 := math.Min(a1, b1)
	fmt.Println("Minimum value is", c1)

	// age := 29      // age is int
	// age = "naveen" // error since we are trying to assign a string to a variable of type int
}
