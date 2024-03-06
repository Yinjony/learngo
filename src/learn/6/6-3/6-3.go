package main //copy与append续
import "fmt" //copy是深拷贝，两个切片数组空间不共享，地址不一样，一般的s:=t,是浅拷贝，共享数组空间

func main() {
	fmt.Println("切片d:")
	d := []int{1, 2, 3}
	d = append(d, 0)   //开辟一个新的位置
	copy(d[2:], d[1:]) //copy语句，去操作（第三个操作中会创建一个临时对象，我们可以借用copy函数避免这个操作，这种方式操作语句虽然冗长了一点，但是相比前面的方法，可以减少中间创建的临时切片。）
	d[1] = 0
	fmt.Println(d)

	fmt.Println("切片e:")
	e := []int{1, 2, 3, 4, 5, 6, 7, 9, 0}
	e = e[2:] //删除开头n个元素
	fmt.Println(e)
	e = append(e[:0], e[2:]...) //append方式
	fmt.Println(e)
	e = e[:copy(e, e[2:])] //copy方式
	fmt.Println(e)

	fmt.Println("切片f:")
	f := []int{1, 2, 3, 4, 5, 6, 7, 9, 0}
	f = append(f[:2], f[4:]...) //append方式删除中间部分
	fmt.Println(f)
	f = f[:copy(f[:2], f[4:])] //copy方式
	fmt.Println(f)

	fmt.Println("切片g:")
	g := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	g = g[:len(g)-2] //删除结尾部分（最快的）
	fmt.Println(g)

	fmt.Println("切片x:")
	x1 := []int{1, 2, 3, 4}
	var x2 []int
	copy(x2, x1)
	fmt.Println(x2) //x2此时还没有长度，所以无法复制
	fmt.Println(x1)
	x2 = make([]int, 2)
	count := copy(x2, x1) //默认从头复制
	fmt.Println(x2)       //x2此时有长度了，但也只是2,长度小，copy函数复制长度以len小的为准
	fmt.Println(x1)
	fmt.Println(count)
	//应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。
}
