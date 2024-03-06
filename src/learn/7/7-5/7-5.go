package main //函数4，函数作为一种实参

import "fmt"

type FormatFunc func(s string, x, y int) string //定义一种函数类型，

func main() {
	s1 := test(func() int { return 100 }) //传递一个函数
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 20) //这里就是将引号内容和几个变量传递到前面的函数中，然后格式化，赋值给s2
	fmt.Println(s1, s2)
}
func test(fn func() int) int { //fn作为一种函数类形参
	return fn()
}
func format(fn FormatFunc, s string, x, y int) string { //fn由FormatFunc定义，是一种函数
	return fn(s, x, y) //意味着传递过来一个函数然后这个函数对其他的变量进行操作
}
