package main

import (
	"fmt"
)

func main() {
	hello_point := "Hello"
	HelloGo.AddWorldPointer(&hello_point)
	fmt.Println(hello_point)

	hello := HelloGo.AddWorld("hello")
	fmt.Println(hello)

}
