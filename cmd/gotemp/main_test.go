package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestRunMultiple(t *testing.T) {
	testArgs := []string{"", "Foo", "Bar\u0007"}
	testWantOuts := []string{"Hello, World!", "Hello, Foo!", "name contains unprintable character(s)"}
	testWantRets := []int{0, 0, 1}

	for i, arg := range testArgs {
		args := os.Args
		stdout := os.Stdout
		stderr := os.Stderr
		os.Args = []string{"gotemp"}
		if arg != "" {
			os.Args = append(os.Args, arg)
		}
		r, w, _ := os.Pipe()
		os.Stdout = w
		os.Stderr = w
		ret := run()
		w.Close()
		bytes, _ := ioutil.ReadAll(r)
		out := strings.TrimSpace(string(bytes))
		os.Args = args
		os.Stdout = stdout
		os.Stderr = stderr

		if out != testWantOuts[i] || ret != testWantRets[i] {
			t.Errorf(`Run with os.Arg[1]="%s" Expected out="%s" return=%d; Got out="%s" return=%d`, arg, testWantOuts[i], testWantRets[i], out, ret)
		}

	}

}
