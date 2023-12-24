/**
 * file: customer-sample.go
 * description: file responsible for explaining the usage of 'fmt.Stringer' interface in Go.
 * data: 12/24/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Customer struct {
	Name string
	Age  int
}

func (customer *Customer) WriteJSON(wj io.Writer) error {
	js, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	_, err = wj.Write(js)
	return err
}

func main() {
	customer := &Customer{
		Name: "Alice",
		Age:  21,
	}

	var buf bytes.Buffer
	err := customer.WriteJSON(&buf)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = customer.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
