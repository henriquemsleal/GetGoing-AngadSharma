package main

import "fmt"

func main() {
	fmt.Println("vim.go")

	// continue, break, switch case,

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	flag := true

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		}
		fmt.Println(i)
	}

	fmt.Println(flag)

	// switch
	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("Default")
	}

	// outra forma
	switch {
	case day == "Fri":
		fmt.Println("TGIF")
		break
	}
}
