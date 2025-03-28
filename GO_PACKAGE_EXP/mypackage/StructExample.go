package mypackage

import (
	"fmt"
)

func StructExample() {
	fmt.Println("Example for struct")

	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "Shashi", Age: 20}

	// Print struct

	fmt.Println("Name", p1.Name)
	fmt.Println("Age", p1.Age)
}
