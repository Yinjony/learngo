package main //map集合
import (
	"fmt"
)

func main() {
	var m1 map[string]string        //初始化1
	m1 = make(map[string]string, 8) //需要进行分配空间，后面的“8”是最大长度，可要可不要

	m2 := make(map[string]string, 8) //初始化2

	m3 := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tjojoyo",
		"India":  "New delhi"}
	//初始化3并赋值
	m2["Franch"] = "Paris" //赋值，这样赋值，里面的元素无序
	m2["Japan"] = "Tjojoyo"

	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)
	fmt.Println(m2["Franch"])
	fmt.Println("m3=", m3)

	fmt.Println("遍历键值对：")
	for k, v := range m3 { //遍历
		fmt.Println(k, v)
	}
	fmt.Println("遍历键：")
	for k := range m3 {
		fmt.Println(k)
		//fmt.Println(m3[k])可以如此顺便遍历值
	}
	fmt.Println("遍历值：")
	for _, k := range m3 {
		fmt.Println(k)
	}
	/*要实现 map 的有序遍历，需要对键进行排序，再按照键的顺序输出值。
	对键进行排序的方法是把所有的键放在一个切片里，然后用sort包中的函数进行排序。*/
	fmt.Println("判断某个键是否存在：")
	n1, jojo := m3["Japan"] //jojo是布尔值（名称可修改哦），有的话n1被赋值
	fmt.Println(n1)
	fmt.Println(jojo)
	n2, jojo := m3["Jojo"] //没有的话，n2为空
	fmt.Println(n2)
	fmt.Println(jojo)

	m1["小明"] = "牛逼"
	m1["小刚"] = "社会"
	m1["小李"] = "逆天"
	fmt.Println("m1=", m1)
	delete(m1, "小李") //删除某个键值对
	fmt.Println(m1)

	fmt.Println("切片集合：")
	var s = make([]map[string]string, 3) //切片集合
	fmt.Println(s)
	s[0] = make(map[string]string, 10) //初始化
	s[0]["name"] = "虎哥"
	s[0]["skill"] = "整活"
	fmt.Println(s)

	fmt.Println("集合切片：")
	var m4 = make(map[string][]string, 3) //集合切片
	fmt.Println(m4)
	m4["小亮"] = []string{"草", "走", "忽略"}
	fmt.Println(m4)
	m4["刀哥"] = append(m4["刀哥"], "fw", "重量级")
	fmt.Println(m4)

}
