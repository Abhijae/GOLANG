package mypackage

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func StructExample() {
	fmt.Println("Example for struct")

	p1 := Person{Name: "Shashi", Age: 20}

	// Print struct

	fmt.Println("Name", p1.Name)

	NameChange(&p1, "Saleem")

	fmt.Println("Name", p1.Name)

	fmt.Println("Age", p1.Age)
}

func NameChange(p *Person, newName string) {
	p.Name = newName
}
