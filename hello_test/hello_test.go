package main

import (
	"testing"
)

func TestHelloEmptyArg(t *testing.T) {

	emptyResult := hello("")

	if emptyResult != "Hello Anonymous!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Anonymous!", emptyResult)
	} else {
		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Anonymous!", emptyResult)
	}

}

func TestHelloValidArg(t *testing.T) {

	jhonResult := hello("Jhon")

	if jhonResult != "Hello Jhon!" {
		t.Errorf("hello(\"Jhon\") failed, expected %v, got %v", "Hello Jhon!", jhonResult)
	} else {
		t.Logf("hello(\"Jhon\") success, expected %v, got %v", "Hello Jhon!", jhonResult)
	}

}
