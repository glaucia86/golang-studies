/**
 * file: echo/main.go
 * description: file responsible for running the main application.
		This application prints its command-line arguments.
 * data: 11/16/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/* Exercise 01: Modify the echo program to also print os.Args[0],
the name of the command that invoked it.
**/

// Line: 18: fmt.Println(strings.Join(os.Args, " "))

/**
Exercise 02: Modify the echo program to print the index and
value of each of its arguments, one per line.

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i+1, arg)
	}
}
**/
