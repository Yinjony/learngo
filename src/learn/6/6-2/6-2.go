package main //append与容量
import "fmt"

func main() {
	fmt.Println("切片a:")
	var a []int            //切片初始化1
	a = append(a, 2, 3, 4) //为切片添加元素
	fmt.Println(a)
	a = append(a, []int{44, 33}...) //添加一个切片
	fmt.Println(a, len(a), cap(a))  //长度与容量（切片最长可达到多少）

	fmt.Println("切片b:")
	var b = []int{1, 2, 3}
	fmt.Println(b)
	b = append([]int{-3, -2, -1}, b...) //在头部添加元素，不过性能要差很多，因为要遍历一遍
	fmt.Println(b)

	fmt.Println("切片c:")
	c := []int{1, 2, 3}
	fmt.Println(c)
	c = append(c[:1], append([]int{6, 7, 8}, c[1:]...)...) //指定位置添加元素（每个添加操作中的第二个append调用都会创建一个临时切片，并将a[i:]的内容复制到新创建的切片中，然后将临时创建的切片再追加到a[:i]。）
	fmt.Println(c)
	c1 := append(c, 23) //返回新的对象，不会改变原来的切片
	fmt.Println(c, c1)

	fmt.Println("cap容量相关：")
	data2 := [...]int{0, 1, 2, 3, 4, 10: 0} //数组
	sx := data2[:2:3]                       //[low:high:max]，新切片的容量为max-low
	fmt.Println(sx)
	fmt.Println(len(sx), cap(sx))
	sx = append(sx, 100, 200, 300) // 一次 append 三个值，超出容量限制。
	fmt.Println(sx, data2)         // 超出后重新分配底层数组，与原数组无关，双方地址不再相同，即切片不再是此数组的引用
	//这里要说明，超出容量后，新切片的容量会扩增2倍（小于1024，若大于则1.25倍），但只要不超过容量便不会有事
}
