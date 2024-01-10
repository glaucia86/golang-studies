/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 01/10/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m) // map[]

	m = map[string][]string{
		"coffee": {"Espresso", "Capuccino", "Mocha"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}

	fmt.Println(m)           // map[coffee:[Espresso Capuccino Mocha] tea:[Hot Tea Chai Tea Chai Latte]]
	fmt.Println(m["coffee"]) // [Espresso Capuccino Mocha]

	m["water"] = []string{"Sparkling"}
	fmt.Println(m) // map[coffee:[Espresso Capuccino Mocha] tea:[Hot Tea Chai Tea Chai Latte] water:[Sparkling]]

	delete(m, "tea")
	fmt.Println(m) // map[coffee:[Espresso Capuccino Mocha] water:[Sparkling]]

	fmt.Println(m["tea"])
	v, ok := m["tea"]
	fmt.Println(v, ok) // [] false

	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}

	fmt.Println(m)
	fmt.Println(m2)
}
