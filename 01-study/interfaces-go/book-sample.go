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

func (book Book) String() string {
	return fmt.Sprintf("Book: %s - %s", book.Title, book.Author)
}

type Count int

func (count Count) String() string {
	return strconv.Itoa(int(count))
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
