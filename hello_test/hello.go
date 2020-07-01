package main

import (
	"fmt"
)

func hello(name string) string {

	if len(name) == 0 {
		return "Hello Anonymous!"
	}
	return fmt.Sprintf("Hello %v!", name)
}
