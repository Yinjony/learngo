package main //循环与选择

import (
	"fmt"
)

func main() {
	//for{}无限循环
	w := 1
	for w > 0 { //相当于while
		fmt.Println(w)
		w--
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("遍历函数")
	s := "hello"
	for i, j := range s { //字符串，但是返回的都是int
		fmt.Println(i, string(j))
	}
	a := [4]int{8, 9, 22, 4}
	for i, j := range a { //数组或切片
		fmt.Println(i, j)
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		a[i] = j + 100 //range会复制对象保留一个副本，所以j中保存的是原来对象的值，但似乎切片是实时变化的
	}
	fmt.Println(a)
	m := map[string]string{"jojo": "hhh", "hu": "diao"}
	for i, j := range m { //map集合
		fmt.Println(i, j)
	}
	//还可以遍历通道，但不方便展示，返回值便是通道里的元素

	b := 0
	switch b { //自带break
	case 0, 2, 3: //可以判断多个值
		fmt.Println("is 0")
		fallthrough //用这个关键字来强行运行下一个case
	case 1:
		fmt.Println("is 1")
	default:
		fmt.Println("gun")
	}

	var c1, c2, c3 chan int
	var i1, i2 int
	c2 = make(chan int, 1)
	i2 = 2
	select { //select语句专门用于通道，系统检查每条case，若能运行，就保留，之后从所有能运行的case中随机选择一条来运行；如果没有，则运行default语句，如果没有default则通道堵塞
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2: //上面的i2=2，又为c2开辟了空间，使得这条语句可以运行
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
