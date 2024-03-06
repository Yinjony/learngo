package main

import "fmt"

func main() {
	a := 100
	b := 200
	fmt.Println("a=", a, "b=", b)
	b, a = a, b
	fmt.Println("a=", a, "b=", b) //多重赋值，变量交换,排序好用
}
