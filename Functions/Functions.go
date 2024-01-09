package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Multiple return values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	price, no := 108, 27
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	// Multiple return values
	area, perimeter := rectProps(24.8, 2.34)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	myarea1, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", myarea1)

}
