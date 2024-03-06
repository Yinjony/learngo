package main //杂

import (
	"fmt"
) //如果是."fmt"，则就省略了fmt这个前置符，如果前面(f"fmt")，那么就是改了别名

const ss = "const" //常量定义
const ab int = 12
const (
	n1 = iota //每加一行常量，iota便加一
	n2
	n3 //如此可以搞枚举
	_  //跳过
	n5
	n6 = 100  //插队
	n7        //不改就还是100
	n8 = iota //回归赋值
	n9
)

func main() {
	fmt.Println(ss, ab)
	fmt.Println(n1, n2, n3, n5, n6, n7, n8, n9)

	type Newint int  //类型定义
	type Myint = int //类型别名
	var a Newint
	var b Myint
	fmt.Printf("\ntype of a:%T\n", a) //表示main包下定义的NewInt类型。这是一个新的数据类型。
	fmt.Printf("type of b:%T\n", b)   //MyInt类型别名只会在代码中存在，编译完成时并不会有MyInt类型。

	s := `jojo
hhhh`
	fmt.Println(s) //多行输出用反引号

	var x *int
	x = new(int) //new函数，为指针分配空间
	*x = 10
	fmt.Println(*x)
}
