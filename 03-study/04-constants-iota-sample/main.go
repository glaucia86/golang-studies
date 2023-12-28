/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 12/28/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

func main() {
	const a = 42
	var i int = a       // 42
	var f64 float64 = a // 42

	fmt.Println(i, f64)

	// iota sample
	const c = iota
	fmt.Println(c)

	const (
		d = 2 * 5
		e             // 2 * 5
		f = iota      // porque está na posição 2
		g             // 3. Porque está na posição 3
		h = 10 * iota // 10 * 4. Por que? Porque pegou o valor da posição 4 (iota) e multiplicou por 10
	)

	fmt.Println(d, e, f, g, h)
}
