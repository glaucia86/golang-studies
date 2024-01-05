/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 01/04/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println(s) // [  ]

	s = []string{"Coffee", "Espresso", "Capuccino"}
	fmt.Println(s) // [Coffee Espresso Capuccino]

	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)

	slices.Delete(s, 1, 2)
	fmt.Println(s)

	// fmt.Println(s[1]) // Espresso
	// s[1] = "Mocha"
	// fmt.Println(s) // [Coffee Mocha Capuccino]

	// s2 := s
	// fmt.Println(s, s2)
}
