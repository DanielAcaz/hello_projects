package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func main() {
	greetMessageEmpty := hello("")
	fmt.Println(aurora.Yellow(greetMessageEmpty))

	greetMessageJhon := hello("Jhon")
	fmt.Println(aurora.Yellow(greetMessageJhon))

}