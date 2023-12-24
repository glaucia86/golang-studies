/**
 * file: book-sample.go
 * description: file responsible for explaining the usage of 'fmt.Stringer' interface in Go.
 * data: 12/24/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

func main() {
	book := Book{
		"Alice in Wonderland",
		"Lewis Carroll",
	}

	WriteLog(book)

	count := Count(3)
	WriteLog(count)
}
