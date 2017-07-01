package main

import "os"
import "fmt"

const pi float64 = 3.14159265358979323846
const zero = 0.0 //无类型浮点型常量
const (
	size int64 = 1024 //
	eof        = -1   //无类型整形常量
)
const u, v float32 = 0, 3   //常量多重赋值
const a, b, c = 3, 4, "foo" //无类型整形 和字符串常量
const mask = 1 << 3

//const Home = os.GetEnv("HOME"常量) 这句话错误，因为getenv在运行的时候才能知道！

//预定义常量 true false iota
// itoa 在下一个const 出现以前 每出现一次 就会加1
const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	a0 = 1 << iota
	b0 //可以简写
	c3
)

const (
	Sunday = iota //包外可见
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberofDays //这个常量没有导出
)

//错误示范

func main() {
	var v1 bool
	v2 := (1 == 2)
	v1 = true
	var PATH = os.Getenv("PATH")
	fmt.Println(PATH)
	fmt.Println(mask)
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(a0)
	fmt.Println(b0)
	fmt.Println(c3)
	fmt.Println(v1)
	fmt.Println(v2)

}
