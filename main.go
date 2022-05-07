package main

import "fmt"

type Person struct {
	name string
	job  string
}

func main() {
	i := 1
	max := 20

	for i < max {
		fmt.Println(i)
		i += 1
	}
	var aperson Person
	aperson.name = "Albert"
	aperson.job = "Developer"

	fmt.Printf("aperson.name = %s\n", aperson.name)
	fmt.Printf("aperson.job = %s\n", aperson.job)
}
