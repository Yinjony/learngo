package main //goroutine与通道（这个与goroutine配合使用，两个并发之间传输数据
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go say("world") //并发进行
	say("hello")
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //转到并发运行
		fmt.Println("hello")
	}

	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() //终止此并发运行
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	fmt.Println("main")

	fmt.Println("通道：")
	str1 := []string{"hello", "world", "!"}
	ch1 := make(chan string, len(str1)) //通道的声明与初始化，可以使int等类型，int值是缓冲区长度，好比你想象的通道长度
	for _, str := range str1 {
		ch1 <- str //先进
	}
	for i := 0; i < len(str1); i++ {
		elem := <-ch1 //先出（想象一下一条通道）
		fmt.Println(elem)
	}
	//单向通道：chan<-,只能发送（是说向通道发送数据）不能接收的通道；<-chan,只能接收不能发送的通道。可以约束输入与输出

	fmt.Println("通道阻塞：")
	/*发送阻塞
	ch1 := make(chan int, 1)
	ch1 <- 1满了
	ch1 <- 2
	接收阻塞
	ch2:=make(chan int,1)
	<-ch2没有接收者
	*/
	ch2 := make(chan string, 0) //没有缓冲区，要同步发送和接收
	go func() {
		for _, str := range str1 {
			ch2 <- str
		}
	}() //所以用了go来并发，发一个收一个，发三次收三次
	for i := 0; i < len(str1); i++ {
		elem := <-ch2
		fmt.Println(elem)
	}
	/*
		var ch3 chan int 没有初始化，通道是个nil
		ch3<-1 阻塞
		<-ch3 阻塞
	*/

	ch4 := make(chan int, 1)
	ch4 <- 1
	close(ch4)        //关闭了通道，无法再向其发送值，但还可以接收值，一般不发送了就用close关了通道
	ele1, ok := <-ch4 //检验是否还有数据留在通道内，有的话就赋给ele1了
	fmt.Println(ele1, ok)
	ele := <-ch4
	ele2, ok := <-ch4 //没有数据了
	fmt.Println(ele2, ok)
	fmt.Println(ele) //被赋了空值
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond) //停止，让并发进行的程序进行
		fmt.Println(s)
	}
}
