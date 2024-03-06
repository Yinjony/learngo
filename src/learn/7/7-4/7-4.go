package main //函数3，匿名变量与闭包
import "fmt"

func main() {
	// --- 匿名函数变量---函数可作为变量
	fn := func() { fmt.Println("Hello, World!") }
	fn()

	// --- 匿名函数切片 ---
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(fns[0](100))
	fmt.Println(fns[1](100))

	// --- 匿名函数结构体---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	fmt.Println(d.fn())

	// --- 匿名函数channel ???---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	fmt.Println((<-fc)())

	//闭包
	fmt.Println("ab:")
	a := add1()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	b := add1()
	fmt.Println(b())
	fmt.Println(b()) //a与b是不同的环境，但是应用的函数相同

	fmt.Println("tmp:")
	tmp1 := add2(10) //传递不同的值，在不同的环境中改变
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add2(100)
	fmt.Println(tmp2(1), tmp2(2))

	fmt.Println("f:")
	f1, f2 := integert(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4)) //由于两个函数在同一个闭包下，base值是一直改变的
}
func add1() func() int { //声明返回值是一个函数，也就是a()嵌套一个b()函数
	i := 0              //同一个闭包下，这个i的值是累计的
	return func() int { //返回一个函数
		i += 1 //闭包复制的是原对象指针，也就是地址
		return i
	}
}

func add2(base int) func(int) int { //可以在声明时设置形参，感觉与add1的i声明没啥区别
	return func(i int) int {
		base += i
		return base
	}
}

func integert(base int) (func(int) int, func(int) int) { //可以返回两个闭包
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
