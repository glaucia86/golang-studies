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
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}

	fmt.Println(s)
}
