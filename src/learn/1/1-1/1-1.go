package main //变量命名

import "fmt"

var (
	x1 int
	x2 int
) //全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。

func GetData() (int, int) {
	return 10, 20
}
func main() {
	fmt.Println("第一部分：")
	var x, y int //变量初始化1

	var a, b int = 1, 2 //变量初始化2

	c, d := 123, "jojo" //变量初始化3，只能用于函数内部，左值没有被定义过

	c1, d := 123, "jojo" //重复定义，但至少有一个新的变量
	var (
		a1 int
		a2 string
	) //批量命名
	a1 = 32
	a2 = "liu"

	var name = "hello" //变量初始化4

	fmt.Println(x, y, a, b, c, c1, d)
	fmt.Println(a1, a2)
	fmt.Println(name)

	fmt.Print("全局变量：", x1, x2)

	fmt.Println("\n第二部分：")
	e, _ := GetData() //匿名变量
	_, f := GetData()
	fmt.Println(e, f)
}
