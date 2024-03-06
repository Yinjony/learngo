package main //结构体

import "fmt"

type Books struct {
	title  string //可以没有变量名称，为匿名
	author string
} //结构体

type person struct {
	name string
	city string
	age  int
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Println(a.name, "会动！")
}

type Dog struct {
	feet    int
	*Animal //嵌套匿名结构体
}

func (d *Dog) wang() {
	fmt.Println(d.name, "会汪汪汪")
}
func main() {
	var book1 Books
	book1.title = "JoJo的奇妙冒险"
	book1.author = "荒木飞吕彦" //初始化1

	book2 := Books{"活着", "余华"} //初始化2

	book3 := Books{author: "平凡的世界", title: "路遥"} //初始化3，键值对

	fmt.Print(book1, book2, book3, "\n", book1.author) //打印

	var book1p *Books
	book1p = &book1 //指针使用
	book2p := &book2
	fmt.Print("\n", *book1p, *book2p, book1p.author) //结构体指针访问也是用"."

	var user struct { //中途匿名
		name string
		age  int
	}
	user.name = "666"
	user.age = 18
	fmt.Println(user)

	p9 := newPerson("NB", "haohaohao", 90) //返回一个指针
	fmt.Println(*p9)

	d1 := &Dog{
		feet: 4,
		Animal: &Animal{ //嵌套的是结构体指针
			name: "niuniu",
		},
	}
	d1.move()
	d1.wang() //这样就可以实现继承
}
func newPerson(name, city string, age int) *person { //构造函数
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
