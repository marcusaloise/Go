package main

import "fmt"

type subjacente int

var x subjacente

var y int

func main() {

	fmt.Printf("%v, %T", x, x)
	fmt.Println("")
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%v, %T", y, y)

}
