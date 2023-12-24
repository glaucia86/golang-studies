/**
 * file: creature-sample.go
 * description: file responsible for explaining the usage of * and & in Go.
 * data: 12/24/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature Creature = Creature{
		Species: "shark",
	}

	fmt.Printf("1)%+v\n", creature) // {Species:shark}
	changeCreature(creature)        // {Species:jellyfish}
	fmt.Printf("3)%+v\n", creature) // should be {Species:jellyfish} but prints {Species:shark}
}

func changeCreature(creature Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2)%+v\n", creature)
}
