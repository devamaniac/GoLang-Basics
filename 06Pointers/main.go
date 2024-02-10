package main

import "fmt"

func main() {
	num := 25
	var p = &num // referencing
	fmt.Println("The value of pointer is:", *p)
	fmt.Println("The address of pointer is:", p)

}
