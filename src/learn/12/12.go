package main

import "fmt"

func main() {
	test()
	testx()
}

func test() {
	defer func() {
		if err := recover(); err != nil { //接收错误指令，recover要在defer函数（注意是函数）内直接调用，否则都返回nil
			println(err.(string)) //打印出来
		}
	}()

	panic("panic error!") //发送错误指令
}
func testx() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic") //栈原理，最后defer先使用，panic也会被顶替
	}()
	panic("test panic")
}
