package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bool
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	// Signed integers
	var a1 int = 50
	b1 := 106
	fmt.Println("value of a1 is", a1, "and b1 is", b1)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b

	// Floating Point Type
	a2, b2 := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a2 + b2
	diff := a2 - b2
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	// Complex types
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// string type
	first := "Harsh"
	last := "Choudhary"
	name := first + " " + last
	fmt.Println("My name is", name)

	// Type Conversion
	i := 10
	var j float64 = float64(i) //this statement will not work without explicit conversion
	fmt.Println("j", j)
}
