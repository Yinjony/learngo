package main //fmt包
import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married) //返回成功扫描的数据个数和遇到的任何错误
	fmt.Printf("name:%s age:%d married:%t\n", name, age, married)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married) //同上，不过是格式化扫描，按着特定格式输入，scan不行
	fmt.Printf("name:%s age:%d married:%t\n", name, age, married)
	var (
		name1 string
		name2 string
	)
	fmt.Scanln(&name1, &name2) //返回值同上，但是只要遇到换行符就结束，不管有没有赋值完scanf同理，但scan不行，必须填完
	fmt.Println(name1, name2)
	//Fprint系列，一般用于写入文件
	//Sprint系列，接收传入的数据并返回一个字符串
}
