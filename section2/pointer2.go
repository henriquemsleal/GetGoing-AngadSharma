package main

import "fmt"

func swap(a1, a2 *int) {
	var temp int
	temp = *a2
	*a2 = *a1
	*a1 = temp
}

func main() {
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
}
