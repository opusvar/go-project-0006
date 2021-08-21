package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	i := 0
	for _, arg := range os.Args[1:] {
		i++
        s += sep + arg
		sep = " "
		fmt.Println(s, "index", i)
	}

    fmt.Println(os.Args[0])
}