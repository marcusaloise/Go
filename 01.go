package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	m := fmt.Sprint(x, " ", y, " ", z)
	fmt.Println(m)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}