/**
 * file: stringer-sample.go
 * description: file responsible for explaining the usage of 'fmt.Stringer' interface in Go.
 * data: 12/24/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

type Animal struct {
	Name string
	Age  uint
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (%d years)", a.Name, a.Age)
}

func main() {
	animal := Animal{
		Name: "Prince",
		Age:  3,
	}

	fmt.Println(animal)
}
