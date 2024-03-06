package main //切片

import (
	"fmt"
)

func main() {
	fmt.Println("切片初始化：")
	var x1 []int            //没有任何定义，所以为nil（没有内存空间）
	x2 := make([]int, 0)    //make函数，第一个数是长度，第二个是容量
	x3 := make([]int, 0, 0) //使用make开辟新的空间，不为nil
	x4 := []int{}           //也开辟了新的空间不为nil
	x5 := []int{1, 2, 3}
	x6 := make([]int, 4, 6)
	arr := [5]int{1, 2, 3, 4, 5}
	x7 := arr[1:4] //利用数组构建切片
	x71 := x7[1:2] //利用切片构建切片
	fmt.Println(x1, x2, x3, x4, x5, x6, x7, x71)

	fmt.Println("切片本质，数组的引用：")
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	fmt.Println(data) //修改了原数组
	p := &s[1]        //访问切片底层数组，并加以修改
	*p += 100
	fmt.Println(s)

	fmt.Println("特殊切片：")
	h := [][]int{ //[]int类型的切片，双切片
		{1, 2, 3}, //初始化值
		{100, 200},
		{11, 22, 33, 44},
	}
	fmt.Println(h)
	j := [5]struct { //结构体数组
		x int
	}{} //未初始化
	k := j //切片
	j[1].x = 10
	k[2].x = 20 //均可访问，也可修改
	fmt.Println(j)

}
