package main //数组

import "fmt"

func main() {
	fmt.Println("第一部分：")
	var a [3]float32
	a[1] = 2.0                               //初始化数组1
	b := [3]float32{2.0, 3.0, 4.0}           //2
	var c = [3]float32{2.0, 3.4, 7.0}        //3
	var d = [...]float32{3.2, 9.2, 7.2, 8.9} //4
	var e = [3]float32{1: 2.0, 2: 7.0}       //指定位置赋值1
	f := [5]int{2: 3, 4: 1}                  //2
	fmt.Println(a, b, c, d, e, f)
	a = b
	fmt.Println(a, b)                     //相同数组之间可赋值
	fmt.Println("a数组长度：", len(a), cap(a)) //两种方式获取长度

	g := [...]struct {
		name string
		age  int
	}{
		{"jojo", 18},
		{"hello", 98},
	}
	fmt.Println(g) //结构体与数组结合

	fmt.Println("第二部分：")
	aa := [2][2]int{{1, 2}, {1, 2}}                      //多维数组初始化1
	var bb [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}} //初始化2
	cc := [...][2]int{{1, 1}, {2, 2}}                    //初始化3
	fmt.Println(aa, bb, cc)

	fmt.Println("第三部分：")
	fmt.Println(Sum(a, 3)) //数组与函数
}
func Sum(a [3]float32, size int) float32 {
	var sum float32
	for i := 0; i < size; i++ {
		sum += a[i]
	}
	return sum
}
