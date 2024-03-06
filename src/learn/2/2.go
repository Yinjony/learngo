package main //指针

import "fmt"

func main() {
	fmt.Println("第一部分：")
	var ip *int
	var fp *float64 //指针的声明1
	a := 20
	b := 3.14
	ip = &a
	fp = &b
	var (
		i *int
		j *int //不初始化也是空指针
	)
	i = nil                     //空指针
	fmt.Println(*ip, *fp, i, j) //指针的使用

	c := 2.3
	p := &c //指针的声明2
	fmt.Println(*p)

	fmt.Println("第二部分：")
	d := new(int) //分配新的内存
	fmt.Println(*d)
	*d = 100
	fmt.Println(*d)

	fmt.Println("第三部分：")
	e := []int{1, 2, 3}
	q := &e              //直接赋值，整个e数组
	qx := &e[1]          //单个元素
	var qtr [3]*int      //数组长度不可省略，此为数组指针
	fmt.Println(*q, *qx) //且只能输出整个数组或单个元素，*p[1]不可使用
	for i := 0; i < 3; i++ {
		qtr[i] = &e[i]
	}
	for i := 0; i < 3; i++ {
		fmt.Print(*qtr[i])
	}

	fmt.Println("\n第四部分：")
	f := 200
	fmt.Println(f)
	modify(&f)
	fmt.Println(f)
}
func modify(x *int) {
	*x = 100
}
