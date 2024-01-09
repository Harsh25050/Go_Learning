package main

import (
	"fmt"
	// "math"
)

func main() {
	// Declaring a constant
	const a = 50
	// a = 89 //reassignment not allowed
	fmt.Println(a)

	// Declaring a group of constants
	const (
		name    = "John"
		age     = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// var a = math.Sqrt(4)   //allowed
	// const b = math.Sqrt(4) //not allowed

	// Boolean Constants
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	// defaultBool = customBool //not allowed
	println(defaultBool)
	println(customBool)

	// Numeric Constants
	const a1 = 5
	var intVar int = a1
	var int32Var int32 = a1
	var float64Var float64 = a1
	var complex64Var complex64 = a1
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// Numeric Expressions
	var a2 = 5.9 / 8
	fmt.Printf("a's type is %T and value is %v", a2, a2)
}
