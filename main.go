package main

import (
	"fmt"
)

func main() {
	tube := NewTuber()

	registar1, _ := NewRegistar(tube)
	registar2, _ := NewRegistar(tube)

	registar1.Listen("ev1", func(payload interface{}) {
		fmt.Printf("Got this %+v", payload)
	})

	registar2.Listen("ev2", func(payload interface{}) {
		fmt.Printf("Got this %+v", payload)
	})

	somePayload := struct {
		f1 string
	}{
		"something something darkside",
	}

	registar1.SayForAll("ev1", somePayload)

}
