package main

import (
	"github.com/ajmclaugh/HelloGo"
	"fmt"
)

func main() {
	hello := "Hello"
	HelloGo.AddWorld(&hello)
	fmt.Println(hello)

}
