package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	// arg := os.Args[1]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

	arg := os.Args[1]

	f, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(f)

	io.Copy(os.Stdout, f)
}
