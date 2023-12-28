/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 12/28/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

func main() {
	a, b := 5, 2

	fmt.Println("Soma: ", a+b)
	fmt.Println("Subtração: ", a-b)
	fmt.Println("Multiplicação: ", a*b)
	fmt.Println("Divisão: ", a/b)
	fmt.Println("Módulo: ", a%b)

	fmt.Println("Conversão", float32(a)/float32(b))

	fmt.Println("Igualdade: ", a == b)
	fmt.Println("Diferença: ", a != b)
	fmt.Println("Maior: ", a > b)
}
