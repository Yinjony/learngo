package main //随机数

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//seed默认种子是1
	rand.Seed(time.Now().Unix())
	a := rand.Intn(1000) //范围0-1000
	fmt.Println(a)
}
