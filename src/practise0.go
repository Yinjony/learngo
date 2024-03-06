package main

import "fmt"

func main() {
	a := 1
	b := 2
	if a > b {
		fmt.Println("jojo")
	} else {
		fmt.Println("hello")
	}

	p := 1
	switch p {
	case 1:
		fmt.Println("hello")
	case 2:
		fmt.Println("jojo")
	default:
		fmt.Println("liu")
	}

	//select?

}
