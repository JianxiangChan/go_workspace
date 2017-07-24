package main

import "fmt"

type test struct {
	Name  string
	Clase string
	Age   int
}

func main() {
	b := test{
		Name:  "Jack",
		Clase: "Great 8",
		Age:   25,
	}

	fmt.Println(b.Age, b.Clase, b.Name)
}
