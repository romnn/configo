package main

import (
	"fmt"

	"github.com/romnnn/configo"
)

func run() string {
	return configo.Shout("This is an example")
}

func main() {
	fmt.Println(run())
}
