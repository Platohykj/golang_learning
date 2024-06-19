package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		if index != 0 {
			fmt.Println("Argument", index, "is", arg)
		}
	}
}
