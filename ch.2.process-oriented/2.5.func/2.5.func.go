package main

import "fmt"
import "mymath"

func main() {

	fmt.Println("func test")

	//2.5.2 函数调用

	c, err := mymath.Add(1, -2)
	fmt.Println(c, err)

	//2.5.3不定参数
	unsure_args(1, 2, 8, 4, 5)

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	myprintf(v1, v2, v3, v4)

	//2.5.5 匿名函数
	func(a, b int, z float64) bool {
		fmt.Println("unamed fucn")
		return a*b < int(z)
	}(1, 2, 5.3)

	f := func(a, b int) int {
		return (a + b)
	}
	fmt.Printf("f = %v, f(1, 4) = %d\r\n", f, f(1, 4))

	var j int = 5

	a := func() func() {
		//这样的好处就是保护了i变量 想修改i变量只有这一个入口
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}() //这里不是很明白 这个括号不是应该加到 上一个大括号结尾？
	a()
	j *= 2
	a()

}

//2.5.3不定参数
func unsure_args(args ...int) { //...type是一个syntactic sugar 实际上等同于一个切片数组
	for _, arg := range args {
		fmt.Println(arg)
	}
	//不定参数的传递 两种方式
	unsure_args1(args...)
	unsure_args1(args[1:]...)
}

func unsure_args1(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myprintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value")
		case string:
			fmt.Println(arg, "is an string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is an unknow value")
		}
	}
}
