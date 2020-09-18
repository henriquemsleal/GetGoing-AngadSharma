package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func main() {
	c := Car{"chevy", 1, 2}
	c1 := Car{
		Name:    "chevy",
		Age:     1,
		ModelNo: 2,
	}
	fmt.Println(c, c1)

}
