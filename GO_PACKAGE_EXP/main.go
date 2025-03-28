package main

import (
	"GO_PACKAGE_EXP/mypackage" // Import using the module name
	"fmt"
)

func main() {
	result := mypackage.Add(10, 20)
	fmt.Println("Sum:", result) // Output: Sum: 30
}
