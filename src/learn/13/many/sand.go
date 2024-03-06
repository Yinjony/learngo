package main

import "fmt"

var _ int64 = s()

func init() {
	fmt.Println("init in sandbox.go") //先运行init函数，再运行main函数
} //单一的源文件可以直接定义init函数，不同源文件的运行顺序是和文件名排序决定，还可以用来多个包连用（就是头文件）
func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}
func main() {
	fmt.Println("main")
}
