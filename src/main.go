package main

import (
	"scanner"
	"fmt"
	"io"
)

func main() {
	g := scanner.NewScanner("test.c")
	for {
		str, err := g.Next()
		fmt.Println(str, err)
		if err == io.EOF {
			break
		}

	}
}
