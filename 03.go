package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true
var s string = fmt.Sprint(x, " ", y, " ", z)

func main() {

	fmt.Println(s)

}