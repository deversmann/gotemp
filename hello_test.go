package gotemp

import (
	"testing"
)

func TestSayHelloSuccess(t *testing.T) {
	names := []string{"Foo\u2620", "Bar\u203D"}
	responses := []string{"Hello, Foo\u2620!", "Hello, Bar\u203D!"}
	for i, name := range names {
		msg, err := SayHello(name)
		if msg != responses[i] || err != nil {
			t.Errorf(`SayHello("%q") = %q, %v want "%q", nil`, name, msg, err, responses[i])
		}
	}
}

func TestSayHelloEmpty(t *testing.T) {
	msg, err := SayHello("")
	if msg != "" || err == nil {
		t.Errorf(`SayHello("") = %q, %v want "", error`, msg, err)
	}
}

func TestSayHelloUnprintable(t *testing.T) {
	names := []string{"Foo\u0007", "Bar\u000b"}
	for _, name := range names {
		msg, err := SayHello(name)
		if msg != "" || err == nil {
			t.Errorf(`SayHello("%q") = %q, %v want "", error`, name, msg, err)
		}
	}
}
