package main

import "fmt"

func main() {
	b1 := true
	var b2 bool

	fmt.Println(b1)
	fmt.Printf("%T\n", b1)
	fmt.Printf("type %T\n value %v\n", b2, b1)
}
