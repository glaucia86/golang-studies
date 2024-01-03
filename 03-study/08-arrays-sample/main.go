/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 01/03/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr) // [  ]

	arr = [3]string{"Coffee", "Espresso", "Capuccino"}
	fmt.Println(arr) // [Coffee Espresso Capuccino]

	fmt.Println(arr[1]) // Espresso
	arr[1] = "Mocha"
	fmt.Println(arr) // [Coffee Mocha Capuccino]

	arr2 := arr
	fmt.Println(arr2) // [Coffee Mocha Capuccino]

}
