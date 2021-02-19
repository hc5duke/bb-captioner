package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	for _, f := range files {
		fmt.Printf("FILE: %s\n", f)
	}
}
