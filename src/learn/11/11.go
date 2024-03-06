package main //接口

import (
	"fmt"
)

type Phone interface { //定义一个接口
	call() //里面是方法
}

type NokiaPhone struct { //结构体1
}

func (nokiaPhone NokiaPhone) call() { //结构体1的方法
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct { //结构体2
}

func (iPhone IPhone) call() { //结构体2的方法
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone         //声明一个接口
	phone = new(NokiaPhone) //
	phone.call()            //通过接口实现方法
	var a1 NokiaPhone
	phone = a1   //赋值
	phone.call() //通过接口实现方法
	phone = new(IPhone)
	phone.call()

}
