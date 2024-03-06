package main //函数5，方法

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

type MyInt int

/*匿名字段*/
type User struct {
	id   int
	name string
}
type Manager struct {
	User  //只有定义没有变量
	title string
}

func main() {
	var c Circle //值
	fmt.Println(c.radius)
	c.radius = 10.00 //赋值
	fmt.Println(c.getArea())
	c.changeRadius(20)
	fmt.Println(c.radius)

	b := &Circle{31.2} //指针
	fmt.Println(b.getArea())
	b.changeRadius(41.0)
	fmt.Println(b.radius) //方法中，接受者为指针类型和值类型的方法，指针变量与值变量均可相互调用

	var m1 MyInt
	m1.Hello()
	m1 = 100

	m := Manager{User{1, "Tom"}, "Administrator"}
	m.P()
	m.User.P() //继承复用能力
	/*方法集规则
	1类型 T 的方法集包含全部 receiver T 方法。
	2类型 *T 的方法集包含全部 receiver T + *T 方法。
	3如类型 S 包含匿名字段 T，则 S 和 *S 的方法集包含 T 方法。
	4如类型 S 包含匿名字段 *T，则 S 和 *S 的方法集包含 T + *T 方法。
	5不管嵌入 T 或 *T，*S 的方法集总是包含 T + *T 方法。*/

	u := User{1, "tom"}
	u.P()
	mValue1 := u.P //换个名字
	mValue1()
	mValue2 := User.V //整个类型换个名字可以值或指针来将V（值)赋值
	mValue2(u)
	mExpreesion := (*User).P //只能通过指针来将P赋值
	mExpreesion(&u)

}
func (c Circle) getArea() float64 { //有一个接受者，这就是方法和函数的区别
	return c.radius * c.radius
}

// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(radius float64) {
	c.radius = radius
}

func (m MyInt) Hello() { //其他类型的也可以作为接受者，但是接口不行
	fmt.Println("Hello,this is a int")
}

func (s *User) P() { //传指针但传值，结果会带个&
	fmt.Println(s)
}
func (s User) V() {
	fmt.Println(s)
}
func (s Manager) P() {
	fmt.Println(s)
}
