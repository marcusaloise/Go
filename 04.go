package main

import "fmt"

type subjacente int

var x subjacente

func main() {

	fmt.Printf("%v, %T", x, x)
	fmt.Println("")
	x = 42
	fmt.Println(x)

}

