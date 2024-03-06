package main //time

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //现在时间
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Day()
	minute := now.Minute()
	second := now.Second()
	fmt.Println("current time:", year, month, day, hour, minute, second)
	w := now.Weekday()
	fmt.Println(w)

	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //同上
	fmt.Println(timestamp1, timestamp2)

	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:01:05.000 Mon Jan")) //格式化，以后面的格式进行输出
	fmt.Println(now.Format("15:04 2006/01/02"))
	later := now.Add(time.Hour) //加时间
	fmt.Println(later)
	before := later.Sub(now) //括号外与内的时间差
	fmt.Println(before)
	fmt.Println(later.Equal(now))  //判断与之前的时间是否相同
	fmt.Println(later.Before(now)) //含义自己看
	fmt.Println(now.After(later))  //含义自己看

	newtimer := time.NewTimer(5 * time.Second) //延时器，五秒后进行
	<-newtimer.C                               //必不可少
	fmt.Println("jojo")

	t := time.NewTicker(2 * time.Second) //计时器，每两秒进行一次
	defer t.Stop()                       //配合使用
	tag := 0
	for {
		tag++
		if tag == 10 {
			break
		}
		<-t.C //配合使用
		fmt.Println("The World")
	}
}
