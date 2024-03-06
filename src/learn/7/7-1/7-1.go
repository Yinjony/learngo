package main //函数1

import "fmt"

func main() {
	fmt.Println("第一部分：")
	a := 1
	b := 2
	c := Add(a, b)
	fmt.Println(c)

	fmt.Println("第二部分：")
	e := []int{1, 2, 3}
	Print("jojo", e...)    //切片要展开（解包）实际上是传多个值
	Print("jojo", 1, 2, 3) //传多个值
	Printx(e)              //interface似乎不用解包

	fmt.Println("第三部分：")
	f, _ := test(1, 2, "jojo") //多返回值，可用"_"忽略返回值
	fmt.Println(f)
	_, g := test(1, 2, "jojo")
	fmt.Println(g)
	testx(test(2, 3, "liu")) //多返回值可直接作为其他函数调用实参

	fmt.Println("第四部分：")
	fmt.Println("Addx=", Addx(5, 6))
}
func Add(a, b int) int { //具名函数
	return a + b
}

func Addx(a, b int) (c int) { //给返回值命名
	c = a + b
	return //这种叫裸返回
	//var c=a+b //声明局部变量会导致遮蔽，要return c来显示返回
	//return c
}

func test(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。可以多返回值，多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func testx(x int, s string) {
	fmt.Println(x, s)
}

func Print(s string, n ...int) { //多值形参,多值部分必须是最后一个（...int）
	fmt.Println(s, n) //可以没有返回值，但前提是不能有返回类型
}

func Printx(n ...interface{}) { //使用interface传递任意类型，十分安全
	fmt.Println(n) //但到了这里就要解包了，不解包就会多出一层
	fmt.Println(n...)
}
