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
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
