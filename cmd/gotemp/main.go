package main

import (
	"fmt"
	"os"

	"github.com/deversmann/gotemp"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	arg := "World"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	retval, err := gotemp.SayHello(arg)
	if err != nil {
		fmt.Println(err)
		return FAILURE
	}
	fmt.Println(retval)
	return SUCCESS
}
