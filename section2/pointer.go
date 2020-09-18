package main

import "fmt"

func main() {
	// pointer
	m1 := 2
	ptr := &m1
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
