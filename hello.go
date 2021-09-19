package gotemp

import (
	"fmt"
	"unicode"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty name")
	}
	for _, c := range name {
		if !unicode.IsPrint(c) {
			return "", fmt.Errorf("name contains unprintable character(s)")
		}
	}
	return fmt.Sprint("Hello, ", name, "!"), nil
}
