/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 01/10/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please, select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small:": 1.50, "medium": 2.00, "large": 2.50}},
		{name: "Tea", prices: map[string]float64{"small:": 1.00, "medium": 1.50, "large": 2.00}},
	}

	fmt.Println(menu)
}
