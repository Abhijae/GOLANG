package main

import (
	"GO_PACKAGE_EXP/mypackage" // Import using the module name
	"fmt"
)

func main() {
	result := mypackage.Add(10, 20)
	fmt.Println("Sum:", result)    // Output: Sum: 30
	fmt.Println("hello:", result)  // Output: Sum: 30
	fmt.Println("HAMMME:", result) // Output: Sum: 30

	//// Defining an array

	arr := [5]int{1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	// using range without indexing

	for _, value := range arr {
		fmt.Println(value)
	}

	// Looping over multi dimensional array

	multiarr := [3][3]int{
		{1, 2, 3},
		{2, 4, 6},
		{3, 6, 9},
	}

	for _, row := range multiarr {
		for _, value := range row {
			fmt.Println(value)
		}
	}

}
