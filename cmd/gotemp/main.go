package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deversmann/gotemp"
)

func main() {
	arg := "World"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	retval, err := gotemp.SayHello(arg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(retval)
}
