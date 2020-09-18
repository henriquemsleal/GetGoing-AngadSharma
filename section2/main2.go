package main

import "fmt"

//arrays

func todo() {

	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}
	fmt.Println(arr2)

	arr2 = append(arr2, "is carlos", "!")

	fmt.Println(arr, arr2)
}
func main() {
	todo()
}
