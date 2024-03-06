package main

import "fmt"

func main() {
	s := []int{2, 34, 5, 6, 56} //遍历切片
	for index, value := range s {
		fmt.Println(index, value)
	}
	fmt.Println(s)
	fmt.Println(s[1:3])

	str := "hello world" //字符串本质也是数组，与切片合用
	s1 := str[0:5]
	s2 := str[6:]
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := []byte(str) //改变字符串内容，先转换为一个切片，中文字符用rune
	fmt.Println(s3)
	s3[6] = 'G'
	s3 = s3[:8]
	s3 = append(s3, '!')
	fmt.Println(s3)  //修改完毕
	str = string(s3) //转换回字符串
	fmt.Println(str)

	s4 := s[1:3:5] //[x:y:z]，y-x长度，z-x容量，注意，容量大小<原切片或数组的长度，不进行标注则默认第二个:后的数字到尾部（数组长度为7则z=7）
	fmt.Println(s4, "len=", len(s4), "cap=", cap(s4))
}
