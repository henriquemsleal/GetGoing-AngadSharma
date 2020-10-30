package main

import "fmt"

func main() {
	// ir else for
	f := true
	flag := &f
	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	arr := []string{"angad", "sharma", "meadlename"}
	for i, value := range arr {
		fmt.Println(i, value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "angad"
	mymap["age"] = 20

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v\n", k, v)
	}
}
