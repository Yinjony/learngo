package main //函数2
import "fmt"

func main() {
	fmt.Println("a函数：", a()) //注意，这里似乎有点变化。。自己运行看吧
	fmt.Println("add函数：", add(1, 2))
	fmt.Println("foo函数：", foo())
}
func a() (i int) {
	defer fmt.Println(i) //defer中直接使用函数和使用func函数效果是不一样的
	i++                  //直接使用一个与不用defer没啥区别
	return
}

func add(x, y int) (z int) {
	defer fmt.Println("defer1:", z) //多个defer按着栈的规则先后调用
	defer func() {                  //但是func就不一样了
		fmt.Println("defer2:", z)
		z += 100
		fmt.Println("defer3", z)
	}()
	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
	//return语句可拆分为z=z+200,return,defer则插在两者之间
}
func foo() int { //add是有具名返回值，这里是匿名返回值
	var i int
	defer func() {
		i++
	}()
	return i //假定返回值变量为 “anony”，上面的返回语句可以拆分成以下过程：anony=i,i++,return
	//所以匿名返回值，return返回的仅仅是一个局部变量，defer操作的也仅是这个局部变量而已，并没有对返回值anony进行修改
}
