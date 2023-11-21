/**
 * file: dup/main.go
 * description: Dup1 prints the text of each line that appears more than
		once in the standard input, preceded by its count.
 * data: 11/20/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		counts[line]++
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
