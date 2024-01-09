package main

import (
	"fmt"
)

func main() {
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)

	// Add Element
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	// Zero value of a map
	employeeSalary["steve"] = 12000

	// Retrieving value for a key from a map
	employeeSalary1 := map[string]int{
		"jay":   12000,
		"rohit": 15000,
		"harsh": 9000,
	}
	employee := "subham"
	salary := employeeSalary1[employee]
	fmt.Println("Salary of", employee, "is", salary)

	// Checking if a key exists
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	// Iterate over all elements in a map
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	// Delete
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)

	// Length of the map
	fmt.Println("length is", len(employeeSalary))

}
