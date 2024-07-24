package main

import (
	"fmt"
	"hello/hello"
	"os"
)

func main() {
	fmt.Println(hello.Hello(os.Args))
}
