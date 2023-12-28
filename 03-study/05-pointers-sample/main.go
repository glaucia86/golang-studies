/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 12/28/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

func main() {
	s := "Hello World!"

	p := &s         // p recebe o endereço de memória da variável s
	fmt.Println(p)  // 0xc000098020 - endereço de memória
	fmt.Println(*p) // Hello World! - valor da variável

	*p = "Hello, Gophers!" // alterando o valor da variável através do ponteiro
	fmt.Println(*p)        // Hello Go! - novo valor da variável

	p = new(string) // p recebe o endereço de memória de uma nova variável do tipo string
	fmt.Println(p)  // 0xc000096050 - endereço de memória
	fmt.Println(*p) // "" - valor da variável
}
